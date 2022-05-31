package cloudtruth

import (
	"context"
	"errors"
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
			"force_delete": { // todo: handle this
				Description: "Whether to allow Terraform to delete the CloudTruth Parameter or not",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTemplateCreate")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	templateName := d.Get("name").(string)
	templateDesc := d.Get("description").(string)
	templateCreate := cloudtruthapi.NewTemplateCreate(templateName)
	if templateDesc != "" {
		templateCreate.SetDescription(templateDesc)
	}
	// Template values can be empty
	templateValue := d.Get("value").(string)
	if templateValue != "" {
		templateCreate.SetBody(templateValue) // This will fail when we attempt to create the template if invalid
	}
	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesCreate(context.Background(),
		*projID).TemplateCreate(*templateCreate).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	// Templates exist at the project level and span all environments
	d.SetId(resp.GetId())
	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTemplateRead")
	var diags diag.Diagnostics
	c := meta.(*cloudTruthClient)
	templateName := d.Get("name").(string)
	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesList(ctx, *projID).Name(templateName).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	// There should be only one template found
	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("found %d templates, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return diags
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTemplateUpdate")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
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
			return diag.FromErr(err)
		}
	}
	return resourceTemplateRead(ctx, d, meta)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTemplateDelete")
	c := meta.(*cloudTruthClient)
	templateID := d.Id()

	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = c.openAPIClient.ProjectsApi.ProjectsTemplatesDestroy(ctx, templateID, *projID).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
