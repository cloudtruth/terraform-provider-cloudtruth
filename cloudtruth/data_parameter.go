package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func dataCloudTruthParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth parameter data source",
		ReadContext: dataCloudTruthParameterRead,
		Schema: map[string]*schema.Schema{
			"environment": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": &schema.Schema{
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
	env := d.Get("environment").(string)
	if env == "" {
		if c.config.Environment != "" {
			env = c.config.Environment
		} else {
			return diag.FromErr(errors.New("the CloudTruth environment must be specified at the provider or resource level"))
		}
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
		*projectID).Environment(env).Name(name).Execute()
	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("error looking up parameter %s: %+v", name, r)))
	}
	if resp.GetCount() > 1 {
		return diag.FromErr(errors.New(fmt.Sprintf("unexpectedly found %d results for parameter %s",
			resp.GetCount(), name)))
	}

	// We know there is only one parameter at this point
	// There will also should only ever be one value per parameter per environment per project
	values := resp.GetResults()[0].GetValues()
	if len(values) > 1 {
		return diag.FromErr(errors.New(fmt.Sprintf("Unexpectedly found %d values for parameter %s",
			len(values), name)))
	}
	resp.GetResults()[0].GetValues()
	for _, v := range values {
		tflog.Debug(ctx, fmt.Sprintf("Found value for %s, lookup env %s, resolved env %s",
			name, env, v.GetEnvironmentName()))
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
		Description: "A CloudTruth parameter data source",
		ReadContext: dataCloudTruthParametersRead,
		Schema: map[string]*schema.Schema{
			"environment_values": {
				Description: "A mapping of environments -> parameter values",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

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

	name := d.Get("name").(string)
	projectID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(), *projectID).Name(name).Execute()
	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("error looking up parameter %s: %+v", name, r)))
	}

	// We should only get back one Parameter
	if resp.GetCount() > 1 {
		return diag.FromErr(errors.New(fmt.Sprintf("unexpectedly found %d results for parameter %s",
			resp.GetCount(), name)))
	}

	// We should now have a value for every environment where the parameter is defined, including environments
	// where the value is inherited from an ancestor environment
	paramRes := resp.GetResults()[0]
	values := paramRes.GetValues()
	valueMap := make(map[string]any)
	for envURL, v := range values {
		// Here we take the key, the environment URI e.g.
		// https://api.cloudtruth.io/api/v1/environments/605de718-4545-4463-852d-288999e5066e/
		// then we parse the ID and use it to fetch the environment's name, we need this for environment
		// parameter values that are inherited, otherwise we cannot attribute the value to the child environment's
		// name
		segments := strings.Split(envURL, "/")
		requestedEnvID := segments[6]
		requestedEnv, err := c.getEnvironmentName(ctx, requestedEnvID)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Debug(ctx, fmt.Sprintf("Found value for %s, lookup env %s, resolved env %s",
			name, *requestedEnv, v.GetEnvironmentName()))
		valueMap[*requestedEnv] = v.GetValue()
	}
	err = d.Set("environment_values", valueMap)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(paramRes.Id)
	return nil
}
