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
				DefaultFunc: schema.EnvDefaultFunc(apiKeyVarName, "UNSET"),
				Description: "CloudTruth API Secret",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(baseURLVarName, defaultBaseURL),
				Description: "The base URL for the CloudTruth API",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"parameter_data_source": dataSourceParameter(),
		},
		ResourcesMap: map[string]*schema.Resource{
			      "project_resource": resourceProject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return configureClient(
		ctx,
		authConfig{
			APIKey:	 d.Get("api_key").(string),
			BaseURL: d.Get("base_url").(string),
		},
	)
}
