package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
)

const (
	// todo: accept both the cli form and the the TF_VAR_cloudtruth_api_key form?
	apiKeyVarName  = "CLOUDTRUTH_API_KEY"
	baseURLVarName = "CLOUDTRUTH_BASE_URL" // Allow env overrides
	defaultBaseURL = "https://api.cloudtruth.io/api/v1"
	defaultHost    = "api.cloudtruth.io"
	defaultScheme  = "https"

	projectVarName     = "CLOUDTRUTH_PROJECT"
	environmentVarName = "CLOUDTRUTH_ENVIRONMENT"
	// todo:
	// use this to set the API request user agent:
	// https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2@v2.4.2/helper/schema#Provider.UserAgent
)

type clientConfig struct {
	APIKey      string
	BaseURL     string
	APIVersion  string
	Project     string
	Environment string
	UserAgent   string
	Host        string
	Scheme      string
}

func configureClient(ctx context.Context, clientConf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	swaggerConfig := cloudtruthapi.NewConfiguration()
	swaggerConfig.Host = clientConf.Host
	swaggerConfig.Scheme = clientConf.Scheme
	/* todo: client debug handling
	if len(config.Verbose) >= 3 {
		ctapi.configuration.Debug = true
	}*/
	swaggerConfig.AddDefaultHeader("Authorization", fmt.Sprintf("Api-Key %s", clientConf.APIKey))
	return &cloudTruthClient{
		client:        http.Client{},
		config:        clientConf,
		swaggerClient: cloudtruthapi.NewAPIClient(swaggerConfig),
	}, nil
}
