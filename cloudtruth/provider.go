package cloudtruth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_cloudtruth_api_key", nil),
			},
			"base_url": { // todo: change to host and scheme if going with Swagger
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(baseURLVarName, defaultBaseURL),
				Description: "The base URL for the CloudTruth API",
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
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	// todo: swap out with Swagger generated client
	return configureClient(
		ctx,
		clientConfig{
			APIKey:  d.Get("api_key").(string),
			BaseURL: d.Get("base_url").(string),
			// todo: update, hardcoded for now
			Host:   "api.cloudtruth.io",
			Scheme: "https",
		},
	)
}
