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
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth project.",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,

		Importer: &schema.ResourceImporter{
			StateContext: projectImportHelper,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"parameter_name_pattern": {
				Description: "A regular expression parameter names must match",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"description": {
				Description: "Description of the project",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"parent": {
				Description: "The parent CloudTruth project, defaults to no parent i.e. a top level project",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
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
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		project, r, err = c.openAPIClient.ProjectsAPI.ProjectsCreate(ctx).ProjectCreate(*projectCreate).Execute()
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

/*
We want to support importing projects by name, thus this function and not what could
otherwise be a "pass through" import
*/
func projectImportHelper(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "entering projectImportHelper")
	defer tflog.Debug(ctx, "exiting projectImportHelper")
	c := meta.(*cloudTruthClient)

	projName, err := c.lookupProject(ctx, d.Id())
	if err != nil {
		return nil, err
	}
	err = d.Set("name", projName)
	if err != nil {
		return nil, err
	}
	d.SetId(d.Id())
	return []*schema.ResourceData{d}, nil
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceProjectRead")
	defer tflog.Debug(ctx, "exiting resourceProjectRead")
	c := meta.(*cloudTruthClient)
	projectName := d.Get("name").(string)
	projectID := d.Id()

	var project *cloudtruthapi.Project
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		project, r, err = c.openAPIClient.ProjectsAPI.ProjectsRetrieve(ctx, projectID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceProjectRead: error reading project %s with ID %s",
				projectName, projectID), r, err)
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

	patchedProject := cloudtruthapi.PatchedProjectUpdate{}
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
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.ProjectsAPI.ProjectsPartialUpdate(ctx,
				projectID).PatchedProjectUpdate(patchedProject).Execute()
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

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsAPI.ProjectsDestroy(ctx, projectID).Execute()
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
