package cloudtruth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerVersion represents the current version of the provider.
// todo: It should be overwritten during the release process.
var providerVersion = "dev"

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeyVarName, nil),
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
				Description: "Provider level environment declaration (overridable)",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(protocolVarName, defaultProtocol),
				Description: "The network protocol, default https",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudtruth_parameter":  dataCloudTruthParameter(),
			"cloudtruth_parameters": dataCloudTruthParameters(),
			//"cloudtruth_project": todo()
			//"cloudtruth_environment": todo()
			//"cloudtruth_template": todo()
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudtruth_project": resourceProject(),
		},
		/*
			ResourcesMap: map[string]*schema.Resource{
				"cloudtruth_environment": todo()
			},
			ResourcesMap: map[string]*schema.Resource{
				"cloudtruth_template": todo()
			},*/
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		return configureClient(
			ctx,
			clientConfig{
				APIKey:      d.Get("api_key").(string),
				Project:     d.Get("project").(string),
				Environment: d.Get("environment").(string),
				Domain:      d.Get("domain").(string),
				Protocol:    d.Get("protocol").(string),
				UserAgent:   provider.UserAgent("terraform-provider-cloudtruth", providerVersion),
			})
	}
	return provider
}
