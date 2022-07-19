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

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Environment.",

		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Environment",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parent": {
				Description: "The parent CloudTruth Environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"force_delete": {
				Description: "Whether to allow Terraform to delete the Environment or not, default to false/disallow",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentCreate")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	envName := d.Get("name").(string)
	envDesc := d.Get("description").(string)
	parentEnv := d.Get("parent").(string)
	parentEnvID, err := c.lookupEnvironment(ctx, parentEnv)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceEnvironmentCreate: %w", err))
	}
	envCreate := cloudtruthapi.NewEnvironmentCreate(envName)
	envCreate.SetParent(fmt.Sprintf("%s/environments/%s/", c.config.BaseURL, *parentEnvID))
	if envDesc != "" {
		envCreate.SetDescription(envDesc)
	}

	var resp *cloudtruthapi.Environment
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsCreate(ctx).EnvironmentCreate(*envCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceEnvironmentCreate: error creating environment %s: %w", envName, err)
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

	envID := resp.GetId()
	d.SetId(envID)
	c.addNewEnvToCaches(envName, envID)
	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentRead")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	envName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedEnvironmentList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsList(ctx).Name(envName).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceEnvironmentRead: error reading environment %s: %w", envName, err)
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
		return diag.FromErr(fmt.Errorf("resourceEnvironmentRead: found %d environments, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentUpdate")
	c := meta.(*cloudTruthClient)
	envID := d.Id()
	envName := d.Get("name").(string)
	envDesc := d.Get("description").(string)
	patchedEnv := cloudtruthapi.PatchedEnvironment{}
	hasChange := false
	if d.HasChange("name") {
		patchedEnv.SetName(envName)
		hasChange = true
	}
	if d.HasChange("description") {
		patchedEnv.SetDescription(envDesc)
		hasChange = true
	}
	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsPartialUpdate(ctx,
				envID).PatchedEnvironment(patchedEnv).Execute()
			if err != nil {
				outErr := fmt.Errorf("resourceEnvironmentUpdate: error updating environment %s: %w", envName, err)
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
	d.SetId(envID)
	return resourceEnvironmentRead(ctx, d, meta)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentDelete")
	c := meta.(*cloudTruthClient)
	envID := d.Id()
	envName := d.Get("name").(string)
	forceDelete := d.Get("force_delete").(bool)
	if !forceDelete {
		return diag.Errorf("resourceEnvironmentDelete: environment %s cannot be deleted unless you set the 'force_delete' property to true",
			envName)
	}

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsDestroy(ctx, envID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceEnvironmentDelete: error destroying environment %s: %w", envName, err)
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

	c.removeEnvFromCaches(envName, envID)
	return nil
}
