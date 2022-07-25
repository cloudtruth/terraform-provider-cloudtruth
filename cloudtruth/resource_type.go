package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

const addRuleRetries = 5

func resourceType() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Type, a user-defined type which can be used to create CloudTruth Parameters",

		CreateContext: resourceTypeCreate,
		ReadContext:   resourceTypeRead,
		UpdateContext: resourceTypeUpdate,
		DeleteContext: resourceTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Type, unique per project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the CloudTruth Type",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"base_type": {
				Description: "The name of the base type, can be a builtin (string|int|boolean) or another custom type",
				Type:        schema.TypeString,
				Required:    true,
			},
			"rule": {
				Description: `The rule(s) describing allowable values for a parameter of this type. Add separate blocks per rule. 
Note that string types support max_len|min_len|regex rules, integets support min|max rules and booleans don't support any rules.'`,
				Type:     schema.TypeList,
				Optional: true, // Optional and disallowed with boolean types
				Elem: &schema.Resource{
					Schema: ruleSchema,
				},
			},
		},
	}
}

func resourceTypeCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeCreate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeCreate := cloudtruthapi.NewParameterTypeCreate(typeName)
	typeDesc := d.Get("description").(string)
	if typeDesc != "" {
		typeCreate.SetDescription(typeDesc)
	}
	baseTypeName := d.Get("base_type").(string)
	baseType := c.lookupType(ctx, baseTypeName)
	if baseType == nil {
		return diag.FromErr(fmt.Errorf("unknown base parameter type %s", baseTypeName))
	}
	typeCreate.SetParent(baseType.GetUrl())
	var typeID string
	var typeCreateResp *cloudtruthapi.ParameterType
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		typeCreateResp, r, err = c.openAPIClient.TypesApi.TypesCreate(ctx).ParameterTypeCreate(*typeCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTypeCreate: error creating type %s: %w", typeName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		typeID = typeCreateResp.GetId()
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	if _, ok := d.GetOk("rule"); ok {
		rules := d.Get("rule").([]any)
		for _, r := range rules {
			err := addRuleToType(ctx, c, typeID, r.(map[string]any))
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(typeID)
	return nil
}

func addRuleToType(ctx context.Context, c *cloudTruthClient, typeID string, rule map[string]any) error {
	tflog.Debug(ctx, "addRuleToType")
	retryCount := 0
	var apiError error
	var createTypeRule cloudtruthapi.ParameterTypeRuleCreate
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(rule["type"].(string))
	if err != nil {
		return err
	}
	createTypeRule.SetType(*typeEnum)
	createTypeRule.SetConstraint(rule["constraint"].(string))

	for retryCount < addRuleRetries {
		_, r, err := c.openAPIClient.TypesApi.TypesRulesCreate(ctx, typeID).ParameterTypeRuleCreate(createTypeRule).Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("addRuleToType: %s", err))
			apiError = err
			retryCount++
		} else {
			apiError = nil
			break
		}
	}
	if apiError != nil {
		return apiError
	}
	return nil
}

func resourceTypeRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeRead")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedParameterTypeList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.TypesApi.TypesList(ctx).NameIexact(typeName).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTypeRead: error reading type %s: %w", typeName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("resourceTypeRead: found %d types, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return nil
}

func resourceTypeUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeUpdate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	hasChange := false
	patchedTypeUpdate := cloudtruthapi.NewPatchedParameterType()
	if d.HasChange("description") {
		typeDesc := d.Get("description").(string)
		patchedTypeUpdate.SetDescription(typeDesc)
		hasChange = true
	}

	typeID := d.Id()
	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.TypesApi.TypesPartialUpdate(ctx, typeID).PatchedParameterType(*patchedTypeUpdate).Execute()
			if err != nil {
				outErr := fmt.Errorf("resourceTypeUpdate: error updating type %s: %w", typeName, err)
				if r.StatusCode >= http.StatusInternalServerError {
					return resource.RetryableError(outErr)
				} else {
					return resource.NonRetryableError(outErr)
				}
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	d.SetId(typeID)
	return resourceTypeRead(ctx, d, meta)
}

func resourceTypeDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeDelete")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.TypesApi.TypesDestroy(ctx, typeID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTypeDelete: error deleting type %s: %w", typeName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(typeID)
	return nil
}
