package cloudtruth

import (
	"context"
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
				Required:    true,
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Required:    true,
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
	// todo: add support for provider level env/project lookup
	env := d.Get("environment").(string)
	project := d.Get("project").(string)
	name := d.Get("name").(string)

	projectID := c.lookupProject(ctx, project)
	if projectID == nil {
		// todo: handle this
	}
	resp, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersList(context.Background(), *projectID).Environment(env).Name(name).Execute()
	if err != nil {
		// todo: handle this
		tflog.Debug(ctx, fmt.Sprintf("+%v", r))
	}
	tflog.Debug(ctx, fmt.Sprintf("+%v", r))
	if resp.GetCount() > 1 {
		// todo: handle this
	}
	// We know there is only one parameter at this point
	// The value is indexed by the environment name but we could have an
	// alias for the specified environment if this value is inherited from a parent env
	param := resp.GetResults()[0]
	for _, v := range param.Values {
		tflog.Debug(ctx, fmt.Sprintf("Found value for %s, lookup env %s, resolved env %s",
			name, env, v.GetEnvironmentName()))
		d.Set("value", v.GetValue())
		d.SetId(v.GetId())
		break
	}
	return nil
}
