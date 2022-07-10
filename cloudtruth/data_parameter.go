package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nav-inc/datetime"
	"net/http"
	"time"
)

func dataCloudTruthParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth parameter data source",
		ReadContext: dataCloudTruthParameterRead,
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
				Description: "The name of the parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The parameter's description",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"external": {
				Description: "Whether or not the value is external, defaults to false",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"value": {
				Description: "The parameter's value",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"as_of": {
				Description: "Retrieve the parameter's historical value 'as of' the specified RFC3333 date, mutually exclusive with 'tag'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag": {
				Description: "Retrieve the parameter's historical value that matches a specific tag, mutually exclusive with 'as_of'",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: %w", err))
	}
	name := d.Get("name").(string)

	// Handle as_of and tag filters
	paramListRequest := c.openAPIClient.ProjectsApi.ProjectsParametersList(ctx, *projID).Environment(*envID).Name(name)
	filteredParamListRequest, err := parseParamListFilters(paramListRequest, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var resp *cloudtruthapi.PaginatedParameterList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		resp, r, err = filteredParamListRequest.Execute()
		if err != nil {
			outErr := fmt.Errorf("dataCloudTruthParameterRead: error looking up parameter %s: %+v", name, r)
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

	if resp.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: expected 1 value for parameter %s, found %d instead",
			name, resp.GetCount()))
	}
	// We know there is only one parameter at this point
	// There should only ever be one value per parameter per environment per project
	results := resp.GetResults()
	param := results[0]
	values := param.GetValues()
	paramID := param.GetId()
	if len(values) > 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: unexpectedly found %d values for parameter %s",
			len(values), name))
	}
	for _, v := range values {
		tflog.Debug(ctx, fmt.Sprintf("dataCloudTruthParameterRead: found value for %s, lookup env %s, resolved env %s",
			name, environment, v.GetEnvironmentName()))
		err := d.Set("value", v.GetValue())
		if err != nil {
			return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: %w", err))
		}
		desc := param.GetDescription()
		err = d.Set("description", desc)
		if err != nil {
			return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: %w", err))
		}
		// We use a composite ID - <PARAMETER_ID>:<PARAMETER_VALUE_ID>
		d.SetId(fmt.Sprintf("%s:%s", paramID, v.GetId()))
	}
	return nil
}

// Return a map of parameter names -> parameter values
func dataCloudTruthParameters() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Parameter data source",
		ReadContext: dataCloudTruthParametersRead,
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
				Description: "Filter for all parameter values defined 'as of' the specified RFC3333 date, mutually exclusive with 'tag'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag": {
				Description: "Filter for parameter values matching a specific tag, mutually exclusive with 'as_of'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parameter_values": {
				Description: "The computed values for all CloudTruth Parameters matching the filter(s) and environment",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthParametersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParametersRead")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: %w", err))
	}

	var paramListRequest cloudtruthapi.ApiProjectsParametersListRequest
	paramsMap := make(map[string]any)
	var pageNum int32 = 0
	for {
		paramListRequest = c.openAPIClient.ProjectsApi.ProjectsParametersList(ctx, *projID).Environment(environment)
		if pageNum > 0 {
			paramListRequest = paramListRequest.Page(pageNum)
		}
		filteredParamListRequest, err := parseParamListFilters(paramListRequest, d)

		var resp *cloudtruthapi.PaginatedParameterList
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
			var r *http.Response
			resp, r, err = filteredParamListRequest.Execute()
			if err != nil {
				outErr := fmt.Errorf("dataCloudTruthParametersRead: error looking up parameters in the %s environment: %+v", environment, r)
				if r.StatusCode >= http.StatusInternalServerError { // A 5xx error
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
		results := resp.GetResults()
		pageNum++

		for _, res := range results {
			paramName := res.GetName()
			// There should only be one value but the result is a map, so we iterate over it
			for _, v := range res.GetValues() {
				paramEnvValue := v.GetValue()
				// We need to exclude parameters that do not have values set in the target environment
				if paramEnvValue != "" {
					paramsMap[paramName] = paramEnvValue
				}
			}
		}
		if resp.GetNext() == "" {
			break
		}
	}

	err = d.Set("parameter_values", paramsMap)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: %w", err))
	}

	/* We use a composite ID for a cloudtruth_parameters result - <PROJECT_ID>:<ENVIRONMENT_ID>
	   The results represent a set of possibly filtered parameters in a given environment and project
	   Ideally, these will be shared across all TF resources in a given state
	   However, we do need to be mindful that a user could pull in the set of parameters more than once
	   within the same state, with possible variations on filters, if that turns out to be an issue
	   then we may need to rethink the ID strategy here */
	d.SetId(fmt.Sprintf("%s:%s", *projID, *envID))
	return nil
}

func parseParamListFilters(paramListRequest cloudtruthapi.ApiProjectsParametersListRequest,
	d *schema.ResourceData) (*cloudtruthapi.ApiProjectsParametersListRequest, error) {
	asOf, tag := "", ""
	var ok bool
	// when used in the context of a cloudtruth_parameter (non data source), these two properties will not exist,
	// therefore we check for unset/nil values first
	if _, ok = d.GetOk("as_of"); ok {
		asOf = d.Get("as_of").(string)
	}
	if _, ok = d.GetOk("tag"); ok {
		tag = d.Get("tag").(string)
	}
	if asOf != "" {
		if tag != "" {
			return nil, fmt.Errorf("parseParamListFilters: 'as_of' and 'tag' cannot both be specified as parameter filters")
		}
		asOfTime, err := datetime.Parse(asOf, time.UTC)
		if err != nil {
			return nil, fmt.Errorf("parseParamListFilters: %w", err)
		}
		paramListRequest = paramListRequest.AsOf(asOfTime)
	}
	if tag != "" {
		paramListRequest = paramListRequest.Tag(tag)
	}
	return &paramListRequest, nil
}
