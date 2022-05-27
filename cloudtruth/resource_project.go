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
			"force_delete": {
				Description: "Whether to allow Terraform to delete the project or not",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectCreate")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
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
	tflog.Debug(ctx, "resourceProjectRead")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
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
	tflog.Debug(ctx, "resourceProjectUpdate")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)

	patchedProject := cloudtruthapi.PatchedProject{}
	hasChange := false
	if d.HasChange("name") {
		patchedProject.SetName(projectName)
		hasChange = true
	}
	if d.HasChange("description") {
		patchedProject.SetDescription(projectDesc)
		hasChange = true
	}
	if hasChange {
		_, r, err := c.openAPIClient.ProjectsApi.ProjectsPartialUpdate(ctx,
			projectID).PatchedProject(patchedProject).Execute()
		if err != nil {
			return diag.FromErr(errors.New(fmt.Sprintf("error updating project %s: %+v", projectName, r)))
		}
	}
	d.SetId(projectID)
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectDelete")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	projectName := d.Get("name")
	r, err := c.openAPIClient.ProjectsApi.ProjectsDestroy(context.Background(), projectID).Execute()
	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("error updating project %s: %+v", projectName, r)))
	}
	return nil
}
