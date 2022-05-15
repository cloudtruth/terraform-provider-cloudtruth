package cloudtruth

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	baseURLVarName = "CLOUDTRUTH_BASE_URL" // Allow env overrides
	defaultBaseURL = "https://api.cloudtruth.io/api/v1"

	// todo: possibly add profile support
	// profile = "CLOUDTRUTH_PROFILE" // todo: add profile support

	// todo:  do we want to be able to override via env vars
	// or should we require explicit references in the HCL?
	project     = "CLOUDTRUTH_PROJECT"
	environment = "CLOUDTRUTH_ENVIRONMENT"
)

type clientConfig struct {
	APIKey     string
	BaseURL    string
	APIVersion string
}

func configureClient(ctx context.Context, clientConf clientConfig) (*cloudTruthClient, diag.Diagnostics) {
	tflog.Debug(ctx, "configureClient")
	return &cloudTruthClient{
		client:       http.Client{},
		clientConfig: clientConf,
	}, nil
}
