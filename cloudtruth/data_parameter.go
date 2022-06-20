package cloudtruth

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nav-inc/datetime"
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
			"value": {
				Description: "The parameter's value",
				Type:        schema.TypeString,
				Computed:    true,
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

	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
		*projID).Environment(*envID).Name(name).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: error looking up parameter %s: %+v", name, r))
	}
	if resp.GetCount() > 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParameterRead: unexpectedly found %d results for parameter %s",
			resp.GetCount(), name))
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
				Description: "Filter for all parameter values defined 'as of' the specified ISO 8601 date, mutually exclusive with 'tag'",
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

	// Handle as_of and tag filters
	paramListRequest := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
		*projID).Environment(environment)
	asOf := d.Get("as_of").(string)
	tag := d.Get("tag").(string)
	if asOf != "" {
		if tag != "" {
			return diag.Errorf("dataCloudTruthParametersRead: 'as_of' and 'tag' cannot both be specified as parameter filters")
		}
		asOfTime, err := datetime.Parse(asOf, time.UTC)
		if err != nil {
			return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: %w", err))
		}
		paramListRequest = paramListRequest.AsOf(asOfTime)
	}
	if tag != "" {
		paramListRequest = paramListRequest.Tag(tag)
	}

	resp, r, err := paramListRequest.Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: error looking up parameters in the %s environment in the %s project: %+v",
			environment, project, r))
	}

	results := resp.GetResults()
	paramsMap := make(map[string]any)
	var pageNum int32 = 1
	for results != nil {
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
		if resp.GetNext() != "" {
			pageNum++
			paramListRequest = c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
				*projID).Environment(environment).Page(pageNum)
			resp, r, err = paramListRequest.Execute()
			if err != nil {
				return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: error looking up parameters in the %s environment in the %s project: %+v",
					environment, project, r))
			}
			results = resp.GetResults()
		} else {
			break
		}
	}

	err = d.Set("parameter_values", paramsMap)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthParametersRead: %w", err))
	}

	// We use a composite ID - <PROJECT_ID>:<ENVIRONMENT_ID>
	// The results represent a set of possibly filtered parameters in a given environment and project
	// Ideally, these will be shared across all TF resources in a given state
	// However, we do need to be mindful that a user could pull in the set of parameters more than once
	// within the same state, with possible variations on filters, if that turns out to be an issue
	// then we may need to rethink the ID strategy here
	d.SetId(fmt.Sprintf("%s:%s", *projID, *envID))
	return nil
}
