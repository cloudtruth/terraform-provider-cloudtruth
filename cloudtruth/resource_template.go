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

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "A CloudTruth Template.",
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Template",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Template",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"value": {
				Description: "The non-evaluated value of the CloudTruth Template, use a Template data source to access the evaluated output in a specific environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"project": {
				Description: "The CloudTruth project where the Template is defined",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTemplateCreate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateCreate: %w", err))
	}

	templateName := d.Get("name").(string)
	templateDesc := d.Get("description").(string)
	templateCreate := cloudtruthapi.NewTemplateCreate(templateName)
	if templateDesc != "" {
		templateCreate.SetDescription(templateDesc)
	}
	templateValue := d.Get("value").(string)
	if templateValue != "" {
		templateCreate.SetBody(templateValue) // This will fail when we attempt to create the template if invalid
	}

	var resp *cloudtruthapi.Template
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesCreate(ctx,
			*projID).TemplateCreate(*templateCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTemplateCreate: error creating template %s: %w", templateName, err)
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

	d.SetId(resp.GetId())
	return nil
}

// todo: handle pagination
func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTemplateRead")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateRead: %w", err))
	}
	templateName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedTemplateList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesList(ctx, *projID).Name(templateName).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTemplateRead: error reading template %s: %w", templateName, err)
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
		return diag.FromErr(fmt.Errorf("resourceTemplateRead: found %d templates, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return nil
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTemplateUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateUpdate: %w", err))
	}
	templateName := d.Get("name").(string)
	templateValue := d.Get("value").(string)
	templateDesc := d.Get("description").(string)
	templateID := d.Id()

	patchedTemplate := cloudtruthapi.PatchedTemplate{}
	hasChange := false
	if d.HasChange("value") {
		patchedTemplate.SetBody(templateValue)
		hasChange = true
	}
	if d.HasChange("description") {
		patchedTemplate.SetDescription(templateDesc)
		hasChange = true
	}
	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesPartialUpdate(ctx, templateID,
				*projID).PatchedTemplate(patchedTemplate).Execute()
			if err != nil {
				outErr := fmt.Errorf("resourceTemplateUpdate: error updating project %s: %w", templateName, err)
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

		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceTemplateUpdate: %w", err))
		}
	}
	return resourceTemplateRead(ctx, d, meta)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTemplateDelete")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateDelete: %w", err))
	}
	templateName := d.Get("name").(string)
	templateID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesDestroy(ctx, templateID, *projID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceTemplateDelete: error destroying template %s: %w", templateName, err)
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
	return nil
}
