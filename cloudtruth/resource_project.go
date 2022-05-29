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
// add force delete
// add parent project support + tests, confirm update support i.e. re-parenting
// data source? import support?
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

	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsCreate(ctx).ProjectCreate(*projectCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.GetId())
	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectRead")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)

	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsList(ctx).Name(projectName).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	// There should be only one project found
	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("found %d projects, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return diags
}

// A project's name and description can be updated
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
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsPartialUpdate(ctx,
			projectID).PatchedProject(patchedProject).Execute()
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(projectID)
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectDelete")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	_, err := c.openAPIClient.ProjectsApi.ProjectsDestroy(context.Background(), projectID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
