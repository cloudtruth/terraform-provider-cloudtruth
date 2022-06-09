package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCloudTruthTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth template data source",
		ReadContext: dataCloudTruthTemplateRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"name": {
				Description: "The name of the template",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": {
				Description: "The template's computed value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthTemplateRead")

	// guaranteed to be set to "default" if not explicitly specified
	templateEnv := d.Get("environment").(string)
	templateEnvID, err := c.lookupEnvironment(ctx, templateEnv)
	if err != nil {
		return diag.FromErr(err)
	}

	project := d.Get("project").(string)
	name := d.Get("name").(string)

	projectID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	templateList, r, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesList(context.Background(),
		*projectID).Environment(*templateEnvID).Name(name).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("error looking up template %s: %+v", name, r))
	}
	if templateList.GetCount() > 1 {
		return diag.FromErr(fmt.Errorf("unexpectedly found %d results for template %s",
			templateList.GetCount(), name))
	}

	// There should only be one template with the specified name per environment per project
	template := templateList.GetResults()[0]
	templateID := template.GetId()

	// We extract its body and request that its referenced parameters/templates be interpolated
	t := template.GetBody()
	tflog.Info(ctx, t)

	templatePrevReq := cloudtruthapi.NewTemplatePreviewCreateRequest(t)
	previewCreate, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatePreviewCreate(ctx, *projectID).
		TemplatePreviewCreateRequest(*templatePrevReq).Execute()
	if err != nil {
		return diag.FromErr(err)
	}
	previewBody := previewCreate.GetBody()
	err = d.Set("value", previewBody)
	if err != nil {
		return diag.FromErr(err)
	}

	// We use a composite ID - <ENV_ID>:<TEMPLATE_ID>
	// This data source represents the evaluation of a template in a given environment (and project)
	d.SetId(fmt.Sprintf("%s:%s", *templateEnvID, templateID))
	return nil
}

// Return a map of template names -> template values
func dataCloudTruthTemplates() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Parameter data source",
		ReadContext: dataCloudTruthTemplatesRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth Project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth Environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"as_of": {
				Description: "Filter for all template values defined 'as of' the specified ISO 8601 date, , mutually exclusive with 'tag'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag": {
				Description: "Filter for template values matching a specific tag, mutually exclusive with 'as_of'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"template_values": {
				Description: "The computed values for all CloudTruth Templates matching the filter(s) and environment",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthTemplatesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
