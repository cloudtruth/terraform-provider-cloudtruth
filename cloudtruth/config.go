package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

const (
	// Allow for CLI style API key env vars, in addition to TF_VAR_var_name approach
	// no default
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	// Optional variables with default values, settable/overridable per Terraform state
	domainVarName = "CLOUDTRUTH_DOMAIN"
	defaultDomain = "api.cloudtruth.io"
	schemeVarName = "CLOUDTRUTH_SCHEME"
	defaultScheme = "https"

	// Optional config variables with no defaults - settable/overridable per Terraform state
	projectVarName     = "CLOUDTRUTH_PROJECT"
	environmentVarName = "CLOUDTRUTH_ENVIRONMENT"
)

// providerVersion represents the current version of the provider.
// todo: It should be overwritten during the release process.
var providerVersion = "dev"

type clientConfig struct {
	APIKey      string
	Project     string
	Environment string
	UserAgent   string
	Domain      string
	Scheme      string
}

func configureClient(ctx context.Context, clientConf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	apiConfig := cloudtruthapi.NewConfiguration()
	apiConfig.Host = clientConf.Domain
	apiConfig.Scheme = clientConf.Scheme
	apiConfig.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", clientConf.APIKey))
	apiConfig.AddDefaultHeader("UserAgent", provider.UserAgent("terraform-provider-cloudtruth", providerVersion))
	return &cloudTruthClient{
		config:        clientConf,
		openAPIClient: cloudtruthapi.NewAPIClient(apiConfig),
	}, nil
}
