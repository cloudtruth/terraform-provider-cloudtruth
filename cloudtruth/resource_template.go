package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Description: "Description of the CloudTruth Parameter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"value": {
				Description: "The value of the CloudTruth Parameter, specific to an Environment (which can be overridden/inherited)",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"project": {
				Description: "The CloudTruth project where the Parameter is defined",
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
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesCreate(ctx,
		*projID).TemplateCreate(*templateCreate).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateCreate: %w", err))
	}
	d.SetId(resp.GetId())
	return nil
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTemplateRead")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateRead: %w", err))
	}

	templateName := d.Get("name").(string)
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesList(ctx, *projID).Name(templateName).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateRead: %w", err))
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
		_, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesPartialUpdate(ctx, templateID,
			*projID).PatchedTemplate(patchedTemplate).Execute()
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
	templateID := d.Id()

	_, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesDestroy(ctx, templateID, *projID).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTemplateDelete: %w", err))
	}
	return nil
}
