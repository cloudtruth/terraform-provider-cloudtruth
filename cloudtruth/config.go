package cloudtruth

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	apiKeyVarName = "CLOUDTRUTH_API_KEY"

	baseURLVarName = "CLOUDTRUTH_BASE_URL" // Allow env overrides
	defaultBaseURL = "https://api.cloudtruth.io/api/v1"

	// todo: possibly add profile support
	// profile = "CLOUDTRUTH_PROFILE" // todo: add profile support

	// Not sure we want these overridden via env vars or specificied in HCL
	project = "CLOUDTRUTH_PROJECT"
	environment = "CLOUDTRUTH_ENVIRONMENT"

)

type authConfig struct {
	APIKey  string
	BaseURL string
}

/* We need 2 headers
-H 'accept: application/json'
-H 'Authorization: Api-Key REDACTED'
*/

func configureClient(ctx context.Context, authConf authConfig) (*http.Client,  diag.Diagnostics) {
	return nil, nil
}
