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

const ruleOperationRetries = 5

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
			"type": {
				Description: "The base type of this custom type, can be a builtin (string|int|boolean) or another custom type",
				Type:        schema.TypeString,
				Required:    true,
			},
			"min": {
				Description: `A rule constraint: the minimum value for integer types, the minimum length for string types.
This value is specified as a string because we need a reliable way to distinguish between default/zero values and unset values.  We use
the empty string value for that purpose.`,
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"min_id": {
				Description: "The internal ID of the min rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"max": {
				Description: `A rule constraint: the maximum value for integer types, the maximum length for string types.
This value is specified as a string because we need a reliable way to distinguish between default/zero values and unset values.  We use
the empty string value for that purpose.`,
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"max_id": {
				Description: "The internal ID of the max rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"regex": {
				Description: "A CloudTruth rule constraint: the regular expression a string type must match, only valid with string types",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"regex_id": {
				Description: "The internal ID of the regex rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceTypeCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTypeCreate")
	defer tflog.Debug(ctx, "exiting resourceTypeCreate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeCreate := cloudtruthapi.NewParameterTypeCreate(typeName)
	typeDesc := d.Get("description").(string)
	if typeDesc != "" {
		typeCreate.SetDescription(typeDesc)
	}
	baseTypeName := d.Get("type").(string)
	baseType := c.lookupType(ctx, baseTypeName)
	if baseType == nil {
		return diag.FromErr(fmt.Errorf("resourceTypeCreate: unknown base type %s", baseTypeName))
	}
	var baseParamType *cloudtruthapi.ParameterType
	if baseTypeName != "" {
		baseParamType = c.lookupType(ctx, baseTypeName)
		if baseParamType == nil {
			return diag.FromErr(fmt.Errorf("unknown parameter type %s", baseTypeName))
		}
	}
	var rules map[string]any
	var err error
	if rules, err = validateAndFetchRules(ctx, c, d, baseParamType); err != nil {
		return diag.FromErr(err)
	}
	fmt.Printf("+%v", rules)

	typeCreate.SetParent(baseType.GetUrl())
	var typeID string
	var typeCreateResp *cloudtruthapi.ParameterType
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		typeCreateResp, r, err = c.openAPIClient.TypesApi.TypesCreate(ctx).ParameterTypeCreate(*typeCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeCreate: error creating type %s", typeName), r, err)
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
	tflog.Debug(ctx, "entering addRuleToType")
	defer tflog.Debug(ctx, "exiting resourceTypeCreate")
	retryCount := 0
	var apiError error
	var createTypeRule cloudtruthapi.ParameterTypeRuleCreate
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(rule["type"].(string))
	if err != nil {
		return err
	}
	createTypeRule.SetType(*typeEnum)
	createTypeRule.SetConstraint(rule["constraint"].(string))

	for retryCount < ruleOperationRetries {
		_, r, err := c.openAPIClient.TypesApi.TypesRulesCreate(ctx, typeID).ParameterTypeRuleCreate(createTypeRule).Execute()
		if r.StatusCode >= http.StatusInternalServerError {
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
	tflog.Debug(ctx, "entering resourceTypeRead")
	defer tflog.Debug(ctx, "exiting resourceTypeCreate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedParameterTypeList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.TypesApi.TypesList(ctx).NameIexact(typeName).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeRead: error reading type %s", typeName), r, err)
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
	tflog.Debug(ctx, "entering resourceTypeUpdate")
	defer tflog.Debug(ctx, "exiting resourceTypeUpdate")
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
				return handleAPIError(fmt.Sprintf("resourceTypeUpdate: error updating type %s", typeName), r, err)
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
	tflog.Debug(ctx, "entering resourceTypeDelete")
	defer tflog.Debug(ctx, "exiting resourceTypeDelete")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.TypesApi.TypesDestroy(ctx, typeID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeDelete: error deleting type %s", typeName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(typeID)
	return nil
}
