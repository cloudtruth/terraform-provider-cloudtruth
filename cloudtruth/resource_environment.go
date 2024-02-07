package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Environment.",

		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: envImportHelper,
		},

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
				Default:     "",
			},
			"parent": {
				Description: "The parent CloudTruth Environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceEnvironmentCreate")
	defer tflog.Debug(ctx, "exiting resourceEnvironmentCreate")
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

	var env *cloudtruthapi.Environment
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		env, r, err = c.openAPIClient.EnvironmentsAPI.EnvironmentsCreate(ctx).EnvironmentCreate(*envCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceEnvironmentCreate: error creating env %s", envName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	envID := env.GetId()
	d.SetId(envID)
	c.addNewEnvToCaches(envName, envID)
	return nil
}

func envImportHelper(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "entering envImportHelper")
	defer tflog.Debug(ctx, "exiting envImportHelper")
	c := meta.(*cloudTruthClient)

	envName, err := c.lookupEnvironment(ctx, d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("name", envName)
	if err != nil {
		return nil, err
	}
	d.SetId(d.Id())
	return []*schema.ResourceData{d}, nil
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceEnvironmentRead")
	defer tflog.Debug(ctx, "exiting resourceEnvironmentRead")
	c := meta.(*cloudTruthClient)
	envName := d.Get("name").(string)
	envID := d.Id()

	var env *cloudtruthapi.Environment
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		env, r, err = c.openAPIClient.EnvironmentsAPI.EnvironmentsRetrieve(ctx, envID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceEnvironmentRead: error reading environment %s with ID %s",
				envName, envID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err := d.Set("name", env.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", env.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	parentURL := env.GetParent()
	parentURLSegments := strings.Split(parentURL, "/")
	parentName, err := c.lookupEnvironment(ctx, parentURLSegments[6])
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("parent", parentName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(env.GetId())
	return nil
}

// resourceEnvironmentUpdate matches the UI and does not support reparenting an environment
func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceEnvironmentUpdate")
	defer tflog.Debug(ctx, "exiting resourceEnvironmentUpdate")
	c := meta.(*cloudTruthClient)
	envID := d.Id()
	envName := d.Get("name").(string)
	envDesc := d.Get("description").(string)
	patchedEnv := cloudtruthapi.PatchedEnvironmentUpdate{}
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
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.EnvironmentsAPI.EnvironmentsPartialUpdate(ctx,
				envID).PatchedEnvironmentUpdate(patchedEnv).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceEnvironmentUpdate: error updating environment %s", envName), r, err)
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
	tflog.Debug(ctx, "entering resourceEnvironmentDelete")
	defer tflog.Debug(ctx, "exiting resourceEnvironmentDelete")
	c := meta.(*cloudTruthClient)
	envID := d.Id()
	envName := d.Get("name").(string)

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.EnvironmentsAPI.EnvironmentsDestroy(ctx, envID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceEnvironmentDelete: error destroying environment %s", envName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	c.removeEnvFromCaches(envName, envID)
	return nil
}
