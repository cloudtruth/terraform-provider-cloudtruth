package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// todo:
// add parent env support + tests, confirm update support i.e. re-parenting
// data source? import support?
func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A Cloudtruth Environment.",

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
	envParent := d.Get("parent").(string)
	envParentID, err := c.lookupEnvironment(ctx, envParent)
	if err != nil {
		return diag.FromErr(err)
	}
	envCreate := cloudtruthapi.NewEnvironmentCreate(envName)
	envCreate.SetParent(fmt.Sprintf("%s/environments/%s/", c.config.BaseURL, *envParentID))
	if envDesc != "" {
		envCreate.SetDescription(envDesc)
	}
	resp, _, err := c.openAPIClient.EnvironmentsApi.EnvironmentsCreate(context.Background()).EnvironmentCreate(*envCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.GetId())
	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentRead")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	envName := d.Get("name").(string)

	resp, _, err := c.openAPIClient.EnvironmentsApi.EnvironmentsList(ctx).Name(envName).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	// There should be only one environment in the results
	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("found %d environments, expcted to find 1", len(res)))
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
		_, _, err := c.openAPIClient.EnvironmentsApi.EnvironmentsPartialUpdate(ctx,
			envID).PatchedEnvironment(patchedEnv).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(envID)
	return resourceEnvironmentRead(ctx, d, meta)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceEnvironmentDelete")
	c := meta.(*cloudTruthClient)
	envID := d.Id()
	envName := d.Get("name")
	forceDelete := d.Get("force_delete").(bool)
	if !forceDelete {
		return diag.Errorf("environment %s cannot be deleted unless you set the 'force_delete' property to be true",
			envName)
	}
	_, err := c.openAPIClient.EnvironmentsApi.EnvironmentsDestroy(context.Background(), envID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
