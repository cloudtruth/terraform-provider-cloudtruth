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
				Default:     "",
			},
			"parent": {
				Description: "The Parent CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
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
	tflog.Debug(ctx, "entering resourceProjectCreate")
	defer tflog.Debug(ctx, "exiting resourceProjectCreate")
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)
	projectCreate := cloudtruthapi.NewProjectCreate(projectName)
	if projectDesc != "" {
		projectCreate.SetDescription(projectDesc)
	}
	projectParent := d.Get("parent").(string)
	if projectParent != "" {
		parent, err := c.getProjectURL(ctx, projectParent)
		if err != nil {
			return diag.FromErr(err)
		}
		projectCreate.SetDependsOn(*parent)
	}

	var project *cloudtruthapi.Project
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		project, r, err = c.openAPIClient.ProjectsApi.ProjectsCreate(ctx).ProjectCreate(*projectCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectCreate: error creating project %s", projectName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	projectID := project.GetId()
	d.SetId(projectID)
	c.addNewProjectToCaches(projectName, projectID)
	return nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceProjectRead")
	defer tflog.Debug(ctx, "exiting resourceProjectRead")
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)
	projectID := d.Id()

	var project *cloudtruthapi.Project
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		project, r, err = c.openAPIClient.ProjectsApi.ProjectsRetrieve(ctx, projectID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectRead: error reading project %s", projectName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(project.GetId())
	// explicitly not checking for/allowing reparenting from the provider because the UI does not allow it
	err := d.Set("name", project.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", project.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// A project's name and description can be updated
func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceProjectUpdate")
	defer tflog.Debug(ctx, "exiting resourceProjectUpdate")
	c := meta.(*cloudTruthClient)
	projectID := d.Id()
	projectName := d.Get("name").(string)
	projectDesc := d.Get("description").(string)

	// force_delete is not a property in the API, it is only a guard rail used by this provider
	patchedProject := cloudtruthapi.PatchedProject{}
	hasChange := false

	// explicitly not checking for/allowing reparenting from the provider because the UI does not allow it
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
	tflog.Debug(ctx, "entering resourceProjectDelete")
	defer tflog.Debug(ctx, "exiting resourceProjectDelete")
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
