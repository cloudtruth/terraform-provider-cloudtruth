package cloudtruth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(), *projectID).Environment(env).Name(name).Execute()
	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("Error looking up parameter %s: %+v", name, r)))
	}
	if resp.GetCount() > 1 {
		return diag.FromErr(errors.New(fmt.Sprintf("Unexpectedly found %d results for parameter %s", resp.GetCount(), name)))
	}

	// We know there is only one parameter at this point
	// There will also should only ever be one value per parameter per environment per project
	param := resp.GetResults()[0]
	if len

	for _, v := range param.Values {
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
