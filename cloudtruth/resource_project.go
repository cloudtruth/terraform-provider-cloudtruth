package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-cloudtruth/pkg/cloudtruthapi"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A Cloudtruth project.",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parent": {
				Description: "The Parent CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceProjectCreate")
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)
	projectCreate := cloudtruthapi.NewProjectCreate(projectName)
	if projectDesc != "" {
		projectCreate.SetDescription(projectDesc)
	}

	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsCreate(ctx).ProjectCreate(*projectCreate).Execute()
	if err != nil {
		return diag.FromErr(errors.New(
			fmt.Sprintf("error creating project %s: %+v", projectName, r)))
	}
	d.SetId(resp.GetId())
	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceProjectRead")
	projectName := d.Get("name").(string)

	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsList(ctx).Name(projectName).Execute()
	if err != nil {
		return diag.FromErr(errors.New(
			fmt.Sprintf("error creating project %s: %+v", projectName, r)))
	}
	// There should be only one project
	// todo: test and confirm this
	d.SetId(resp.GetResults()[0].GetId())
	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return diag.Errorf("not implemented")
}
