package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/hashstructure"
	"strconv"
)

// todo:
// possible support for wrap, maybe also evaluate and mask_secrets
// some more tests, including negative cases
// add filters to the plural data type
// add type support or just interpret all param sources as strings?
// the API returns them as strings
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

// FIX: likely need to use the composite ID here too
func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")

	// guaranteed to be set to "default" if not explicitly specified
	paramEnv := d.Get("environment").(string)
	paramEnvID, err := c.lookupEnvironment(ctx, paramEnv)
	if err != nil {
		return diag.FromErr(err)
	}

	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	name := d.Get("name").(string)

	projectID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
		*projectID).Environment(*paramEnvID).Name(name).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("error looking up parameter %s: %+v", name, r))
	}
	if resp.GetCount() > 1 {
		return diag.FromErr(fmt.Errorf("unexpectedly found %d results for parameter %s",
			resp.GetCount(), name))
	}

	// We know there is only one parameter at this point
	// There will also should only ever be one value per parameter per environment per project
	values := resp.GetResults()[0].GetValues()
	if len(values) > 1 {
		return diag.FromErr(fmt.Errorf("Unexpectedly found %d values for parameter %s",
			len(values), name))
	}
	resp.GetResults()[0].GetValues()
	for _, v := range values {
		tflog.Debug(ctx, fmt.Sprintf("Found value for %s, lookup env %s, resolved env %s",
			name, paramEnv, v.GetEnvironmentName()))
		err := d.Set("value", v.GetValue())
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(v.GetId())
	}
	return nil
}

// Return a map of env -> parameter values
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
			"parameter_values": {
				Description: "The computed values for the CloudTruth Parameter",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

// todo: maybe break this up into 2 or more smaller functions
func dataCloudTruthParametersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParametersRead")
	project := d.Get("project").(string)
	if project == "" {
		if c.config.Project != "" {
			project = c.config.Project
		} else {
			return diag.FromErr(errors.New("the CloudTruth project must be specified at the provider or resource level"))
		}
	}
	projectID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}

	// guaranteed to be set to "default" if not explicitly specified
	environment := d.Get("environment").(string)
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
		*projectID).Environment(environment).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("error looking up parameters in the %s environment in the %s project: %+v",
			environment, project, r))
	}

	valueMap := make(map[string]any)
	results := resp.GetResults()
	var pageNum int32 = 1
	for results != nil {
		for _, res := range results {
			paramName := res.GetName()
			// There should only ever be one value but the result is a map, so we iterate over it
			for _, v := range res.GetValues() {
				paramEnvValue := v.GetValue()
				// We need to exclude parameters that do not have values set in the target environment
				if paramEnvValue != "" {
					valueMap[paramName] = paramEnvValue
				}
			}
		}
		// HasNext() doesn't do what we want :(
		if resp.GetNext() != "" {
			pageNum++
			resp, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(),
				*projectID).Environment(environment).Page(pageNum).Execute()
			if err != nil {
				return diag.FromErr(fmt.Errorf("error looking up parameters in the %s environment in the %s project: %+v",
					environment, project, r))
			}
			results = resp.GetResults()
		} else {
			break
		}
	}
	err = d.Set("parameter_values", valueMap)
	if err != nil {
		return diag.FromErr(err)
	}
	hash, err := hashstructure.Hash(valueMap, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// todo: add tests where the value actually changes
	// We hash the contents of the map to determine if any parameters have changed
	// NOTE: this is stable in regards to map order, see
	// https://github.com/mitchellh/hashstructure/blob/master/hashstructure.go#L242
	d.SetId(strconv.FormatUint(hash, 10))
	return nil
}
