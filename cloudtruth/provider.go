package cloudtruth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	p := schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_cloudtruth_api_key", nil),
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(domainVarName, defaultDomain),
				Description: "The CloudTruth API domain name",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(projectVarName, nil),
				Description: "Provider level project declaration (overridable)",
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(environmentVarName, nil),
				Description: "Provider level enivironment declaration (overridable)",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudtruth_parameter": dataCloudTruthParameter(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_project": resourceProject(),
		},
		/*ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_parameters": todo()
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_environment": todo()
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_template": todo()
		},*/
		ConfigureContextFunc: providerConfigure,
	}
	return &p
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	return configureClient(
		ctx,
		clientConfig{
			APIKey:      d.Get("api_key").(string),
			Project:     d.Get("project").(string),
			Environment: d.Get("environment").(string),
			Domain:      d.Get("domain").(string),
			Scheme:      d.Get("scheme").(string),
		},
	)
}
