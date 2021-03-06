package cloudtruth

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testProviderFactories = map[string]func() (*schema.Provider, error){
		"cloudtruth": func() (*schema.Provider, error) {
			return New("dev")(), nil
		},
	}
)

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = New("dev")
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv(apiKeyVarName); v == "" {
		t.Fatalf("%s must be set for acceptance tests", apiKeyVarName)
	}
}
