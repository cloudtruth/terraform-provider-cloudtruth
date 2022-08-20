package cloudtruth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
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
				"cloudtruth_parameter_value":  dataCloudTruthParameterValue(),
				"cloudtruth_parameter_values": dataCloudTruthParameterValues(),
				"cloudtruth_tag":              dataCloudTruthTag(),
				"cloudtruth_template":         dataCloudTruthTemplate(),
				"cloudtruth_templates":        dataCloudTruthTemplates(),
				"cloudtruth_user":             dataCloudTruthUser(),
				"cloudtruth_users":            dataCloudTruthUsers(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"cloudtruth_aws_import_action": resourceAWSImportAction(),
				"cloudtruth_aws_push_action":   resourceAWSPushAction(),
				"cloudtruth_access_grant":      resourceAccessGrant(),
				"cloudtruth_environment":       resourceEnvironment(),
				"cloudtruth_group":             resourceGroup(),
				"cloudtruth_parameter":         resourceParameter(),
				"cloudtruth_parameter_value":   resourceParameterValue(),
				"cloudtruth_project":           resourceProject(),
				"cloudtruth_tag":               resourceTag(),
				"cloudtruth_template":          resourceTemplate(),
				"cloudtruth_type":              resourceType(),
			},
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
					UserAgent:   provider.UserAgent("terraform-provider-cloudtruth", version),
				})
		}
		return provider
	}
}
