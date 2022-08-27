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

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth project.",

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
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)
	projectCreate := cloudtruthapi.NewProjectCreate(projectName)
	if projectDesc != "" {
		projectCreate.SetDescription(projectDesc)
	}

	var resp *cloudtruthapi.Project
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.ProjectsApi.ProjectsCreate(ctx).ProjectCreate(*projectCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectCreate: error creating project %s", projectName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	projectID := resp.GetId()
	d.SetId(projectID)
	c.addNewProjectToCaches(projectName, projectID)
	return nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectRead")
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedProjectList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.ProjectsApi.ProjectsList(ctx).Name(projectName).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectRead: error reading project %s", projectName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// There should be only one project found
	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("resourceProjectRead: found %d projects, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return nil
}

// A project's name and description can be updated
func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectUpdate")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)

	// force_delete is not a property in the API, it is only a guard rail used by this provider
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
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.ProjectsApi.ProjectsPartialUpdate(ctx,
				projectID).PatchedProject(patchedProject).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceProjectUpdate: error updating project %s", projectName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	d.SetId(projectID)
	return resourceProjectRead(ctx, d, meta)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceProjectDelete")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	projectName := d.Get("name").(string)
	forceDelete := d.Get("force_delete").(bool)
	if !forceDelete {
		return diag.Errorf("resourceProjectDelete: project %s cannot be deleted unless you set the 'force_delete' property to be true",
			projectName)
	}

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsDestroy(ctx, projectID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectDelete: error destroying project %s", projectName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	c.removeProjectFromCaches(projectName, projectID)
	return nil
}
