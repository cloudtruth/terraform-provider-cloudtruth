package cloudtruth

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	fmt.Print("help")
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
			"cloudtruth_parameter": dataCloudTruthParameter(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_project": resourceProject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return configureClient(
		ctx,
		authConfig{
			APIKey:  d.Get("api_key").(string),
			BaseURL: d.Get("base_url").(string),
		},
	)
}
