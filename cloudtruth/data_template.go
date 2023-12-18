package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nav-inc/datetime"
	"net/http"
	"time"
)

func dataCloudTruthTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth template data source. Use this to reference the rendered values from CloudTruth
Templates in the context of a project and environment.`,
		ReadContext: dataCloudTruthTemplateRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true, // But must be set via this property or the CLOUDTRUTH_PROJECT env variable
			},
			"environment": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"as_of": {
				Description: "Retrieve the template's historical value 'as of' the specified RFC3333 date, mutually exclusive with 'tag'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag": {
				Description: "Retrieve the template's historical value that matches a specific tag, mutually exclusive with 'as_of'",
				Type:        schema.TypeString,
				Optional:    true,
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

// This data source represents the evaluation of a template in a given environment (and project)
func dataCloudTruthTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering dataCloudTruthTemplateRead")
	defer tflog.Debug(ctx, "exiting dataCloudTruthTemplateRead")
	c := meta.(*cloudTruthClient)
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: %w", err))
	}
	name := d.Get("name").(string)

	templateListRequest := c.openAPIClient.ProjectsAPI.ProjectsTemplatesList(ctx,
		*projID).Environment(*envID).Name(name)
	filteredTemplateListRequest, err := parseTemplateListFilters(templateListRequest, d)
	if err != nil {
		return diag.FromErr(err)
	}
	var templateList *cloudtruthapi.PaginatedTemplateList
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		templateList, r, err = filteredTemplateListRequest.Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("dataCloudTruthTemplatesRead: error reading template template %s", name), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	if templateList.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplateRead: expected 1 value for template %s, found %d instead",
			name, templateList.GetCount()))
	}

	// There should only be one template with the specified name per environment per project
	template := templateList.GetResults()[0]
	templateID := template.GetId()

	// We extract its body and request that its referenced parameters/templates be interpolated
	body := template.GetBody()
	tflog.Debug(ctx, fmt.Sprintf("dataCloudTruthTemplateRead: template body - %s", body))

	previewBody, err := renderTemplateBody(ctx, name, body, *projID, *envID, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("value", previewBody)
	if err != nil {
		return diag.FromErr(err)
	}
	// We use a composite ID - <ENV_ID>:<TEMPLATE_ID>
	d.SetId(fmt.Sprintf("%s:%s", *envID, templateID))
	return nil
}

func renderTemplateBody(ctx context.Context, name, body, projectID, envID string, meta any) (*string, error) {
	tflog.Debug(ctx, "entering renderTemplateBody")
	defer tflog.Debug(ctx, "exiting renderTemplateBody")
	c := meta.(*cloudTruthClient)
	templatePrevReq := cloudtruthapi.NewTemplatePreviewCreateRequest(body)
	var previewBody string
	retryCount := 0
	var apiError error
	// We cannot use the TF Provider SDK's retry functionality because it only works with state change events
	// so we employ a simple retry loop instead
	for retryCount < loadCacheRetries {
		previewCreate, r, err := c.openAPIClient.ProjectsAPI.ProjectsTemplatePreviewCreate(ctx, projectID).Environment(envID).
			TemplatePreviewCreateRequest(*templatePrevReq).Execute()
		if r.StatusCode >= 500 {
			apiError = err
			tflog.Debug(ctx, fmt.Sprintf("renderTemplateBody: error rendering template %s in the %s project: %+v", name,
				projectID, r))
		} else {
			previewBody = previewCreate.GetBody()
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf("renderTemplateBody: error fetching body for template %s in the %s project: %+v", name,
					projectID, err))
			}
			apiError = nil
			break
		}
		retryCount++
	}
	if apiError != nil {
		return nil, fmt.Errorf("renderTemplateBody: %w", apiError)
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
				Description: "Filter for all template values defined 'as of' the specified RFC3333 date, , mutually exclusive with 'tag'",
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
	tflog.Debug(ctx, "entering dataCloudTruthTemplatesRead")
	defer tflog.Debug(ctx, "exiting dataCloudTruthTemplatesRead")
	c := meta.(*cloudTruthClient)
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
	}

	var templateListRequest cloudtruthapi.ApiProjectsTemplatesListRequest
	templateMap := make(map[string]any)
	var pageNum int32 = 0
	for {
		templateListRequest = c.openAPIClient.ProjectsAPI.ProjectsTemplatesList(ctx, *projID).Environment(environment)
		if pageNum > 0 {
			templateListRequest = templateListRequest.Page(pageNum)
		}
		filteredTemplateListRequest, err := parseTemplateListFilters(templateListRequest, d)

		var templateList *cloudtruthapi.PaginatedTemplateList
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
			var r *http.Response
			templateList, r, err = filteredTemplateListRequest.Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("dataCloudTruthTemplatesRead: error reading templates in the %s environment", environment), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
		results := templateList.GetResults()
		pageNum++

		for _, res := range results {
			templateName := res.GetName()
			templateBody := res.GetBody()
			previewBody, err := renderTemplateBody(ctx, templateName, templateBody, *projID, *envID, meta)
			if err != nil {
				return diag.FromErr(fmt.Errorf("dataCloudTruthTemplatesRead: %w", err))
			} else {
				templateMap[templateName] = previewBody
			}
		}
		if templateList.GetNext() == "" {
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

func parseTemplateListFilters(templateListRequest cloudtruthapi.ApiProjectsTemplatesListRequest,
	d *schema.ResourceData) (*cloudtruthapi.ApiProjectsTemplatesListRequest, error) {
	asOf, tag := "", ""
	var ok bool
	// when used in the context of a cloudtruth_template (non data source), these two properties will not exist,
	// therefore we check for unset/nil values first
	if _, ok = d.GetOk("as_of"); ok {
		asOf = d.Get("as_of").(string)
	}
	if _, ok = d.GetOk("tag"); ok {
		tag = d.Get("tag").(string)
	}
	if asOf != "" {
		if tag != "" {
			return nil, fmt.Errorf("dataCloudTruthParametersRead: 'as_of' and 'tag' cannot both be specified as parameter filters")
		}
		asOfTime, err := datetime.Parse(asOf, time.UTC)
		if err != nil {
			return nil, fmt.Errorf("dataCloudTruthParametersRead: %w", err)
		}
		templateListRequest = templateListRequest.AsOf(asOfTime)
	}
	if tag != "" {
		templateListRequest = templateListRequest.Tag(tag)
	}
	return &templateListRequest, nil
}
