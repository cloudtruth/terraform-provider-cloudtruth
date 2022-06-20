package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nav-inc/datetime"
	"time"
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
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: %w", err))
	}
	name := d.Get("name").(string)

	templateList, r, err := c.openAPIClient.ProjectsApi.ProjectsTemplatesList(context.Background(),
		*projID).Environment(*envID).Name(name).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: error looking up template %s: %+v", name, r))
	}
	if templateList.GetCount() > 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: unexpectedly found %d results for template %s",
			templateList.GetCount(), name))
	}

	// There should only be one template with the specified name per environment per project
	template := templateList.GetResults()[0]
	templateID := template.GetId()

	// We extract its body and request that its referenced parameters/templates be interpolated
	body := template.GetBody()
	tflog.Debug(ctx, fmt.Sprintf("dataCloudTruthTemplateRead: template body - %s", body))

	previewBody, err := renderTemplateBody(ctx, body, *projID, meta)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: %w", err))
	}
	_ = d.Set("value", previewBody)
	// We use a composite ID - <ENV_ID>:<TEMPLATE_ID>
	// This data source represents the evaluation of a template in a given environment (and project)
	d.SetId(fmt.Sprintf("%s:%s", *envID, templateID))
	return nil
}

func renderTemplateBody(ctx context.Context, body, projectID string, meta any) (*string, error) {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "renderTemplateBody")
	templatePrevReq := cloudtruthapi.NewTemplatePreviewCreateRequest(body)
	previewCreate, _, err := c.openAPIClient.ProjectsApi.ProjectsTemplatePreviewCreate(ctx, projectID).
		TemplatePreviewCreateRequest(*templatePrevReq).Execute()
	if err != nil {
		return nil, err
	}
	previewBody := previewCreate.GetBody()
	if err != nil {
		return nil, err
	}
	return &previewBody, nil
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
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthTemplatesRead")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
	}

	// Handle as_of and tag filters
	templateListRequest := c.openAPIClient.ProjectsApi.ProjectsTemplatesList(context.Background(),
		*projID).Environment(environment)
	asOf := d.Get("as_of").(string)
	tag := d.Get("tag").(string)
	if asOf != "" {
		if tag != "" {
			return diag.Errorf("dataCloudTruthTemplatesRead: 'as_of' and 'tag' cannot both be specified as parameter filters")
		}
		asOfTime, err := datetime.Parse(asOf, time.UTC)
		if err != nil {
			return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
		}
		templateListRequest = templateListRequest.AsOf(asOfTime)
	}
	if tag != "" {
		templateListRequest = templateListRequest.Tag(tag)
	}

	resp, r, err := templateListRequest.Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: error looking up parameters in the %s environment in the %s project: %+v",
			environment, project, r))
	}

	templateMap := make(map[string]any)
	results := resp.GetResults()
	var pageNum int32 = 1
	for results != nil {
		for _, res := range results {
			templateName := res.GetName()
			templateBody := res.GetBody()
			previewBody, err := renderTemplateBody(ctx, templateBody, *projID, meta)
			if err != nil {
				return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
			} else {
				templateMap[templateName] = previewBody
			}
		}
		if resp.GetNext() != "" {
			pageNum++
			templateListRequest = c.openAPIClient.ProjectsApi.ProjectsTemplatesList(context.Background(),
				*projID).Environment(environment).Page(pageNum)
			resp, r, err = templateListRequest.Execute()
			if err != nil {
				return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: error looking up parameters in the %s environment in the %s project: %+v",
					environment, project, r))
			}
			results = resp.GetResults()
		} else {
			break
		}
	}

	err = d.Set("template_values", templateMap)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
	}

	// As with the parameters (plural) data source,we may need to use a more complex ID strategy here
	d.SetId(fmt.Sprintf("%s:%s", *projID, *envID))
	return nil
}
