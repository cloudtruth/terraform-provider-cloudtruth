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

// todo: add handling for rules
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
			}, /*
				"base_type": {
					Description: "The base type for this custom type",
					Type:        schema.TypeString,
					Required:    true,
				},*/
			"description": {
				Description: "Description of the CloudTruth Type",
				Type:        schema.TypeString,
				Optional:    true,
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
	d.SetId(typeID)
	return nil
}

// todo: handle pagination
func resourceTypeRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTypeRead")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedParameterTypeList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.TypesApi.TypesList(ctx).Execute()
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
