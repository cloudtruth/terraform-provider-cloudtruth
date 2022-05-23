package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	// Allow for CloudTruth CLI API key definition in addition to TF_VAR_var_name approach
	// no default
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	// Optional variables with default values, settable/overridable per Terraform state
	domainVarName   = "CLOUDTRUTH_DOMAIN"
	defaultDomain   = "api.cloudtruth.io"
	protocolVarName = "CLOUDTRUTH_PROTOCOL"
	defaultProtocol = "https"

	// Optional config variables with no defaults - settable/overridable per Terraform state
	projectVarName     = "CLOUDTRUTH_PROJECT"
	environmentVarName = "CLOUDTRUTH_ENVIRONMENT"
)

type clientConfig struct {
	APIKey      string
	Project     string
	Environment string
	UserAgent   string
	Domain      string
	Protocol    string
}

func configureClient(ctx context.Context, conf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	tflog.Debug(ctx, fmt.Sprintf("%+v", conf))
	apiConfig := cloudtruthapi.NewConfiguration()
	apiConfig.Host = conf.Domain
	apiConfig.Scheme = conf.Protocol
	apiConfig.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", conf.APIKey))
	apiConfig.AddDefaultHeader("UserAgent", conf.UserAgent)
	client := cloudTruthClient{
		config:        conf,
		openAPIClient: cloudtruthapi.NewAPIClient(apiConfig),
	}

	// load caches
	err := client.loadProjectNameCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadProjectIDCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadEnvNameCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	err = client.loadEnvIDCache(ctx)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return &client, nil
}
