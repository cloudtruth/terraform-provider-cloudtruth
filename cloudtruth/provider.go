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
				Required:    true,
				Description: "CloudTruth API Secret",
				Sensitive:   true,
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(baseURLVarName, defaultBaseURL),
				Description: "The base URL for the CloudTruth API",
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
	return configureClient(
		ctx,
		clientConfig{
			APIKey:  d.Get("api_key").(string),
			BaseURL: d.Get("base_url").(string),
		},
	)
}
