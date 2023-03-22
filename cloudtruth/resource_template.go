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

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "A CloudTruth Template.",
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,

		Importer: &schema.ResourceImporter{
			StateContext: paramOrTemplateImportHelper,
		},

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
				Default:     "",
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
				Optional:    true, // But must be set via this property or the CLOUDTRUTH_PROJECT env variable
			},
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTemplateCreate")
	defer tflog.Debug(ctx, "exiting resourceTemplateCreate")
	c := meta.(*cloudTruthClient)
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

	var template *cloudtruthapi.Template
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		template, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesCreate(ctx,
			*projID).TemplateCreate(*templateCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTemplateCreate: error creating template %s", templateName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(template.GetId())
	return nil
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTemplateRead")
	defer tflog.Debug(ctx, "exiting resourceTemplateRead")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	templateName := d.Get("name").(string)
	templateID := d.Id()

	var template *cloudtruthapi.Template
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		template, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesRetrieve(ctx, templateID, *projID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTemplateRead: error reading template %s", templateName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err = d.Set("name", template.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", template.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("value", template.GetBody())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(template.GetId())
	return nil
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTemplateUpdate")
	defer tflog.Debug(ctx, "exiting resourceTemplateUpdate")
	c := meta.(*cloudTruthClient)
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
	if d.HasChange("name") {
		patchedTemplate.SetName(templateName)
		hasChange = true
	}
	if d.HasChange("value") {
		patchedTemplate.SetBody(templateValue)
		hasChange = true
	}
	if d.HasChange("description") {
		patchedTemplate.SetDescription(templateDesc)
		hasChange = true
	}
	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesPartialUpdate(ctx, templateID,
				*projID).PatchedTemplate(patchedTemplate).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceTemplateUpdate: error updating template %s", templateName), r, err)
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
	tflog.Debug(ctx, "entering resourceTemplateDelete")
	defer tflog.Debug(ctx, "exiting resourceTemplateDelete")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateDelete: %w", err))
	}
	templateName := d.Get("name").(string)
	templateID := d.Id()

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesDestroy(ctx, templateID, *projID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTemplateDelete: error destroying template %s", templateName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
