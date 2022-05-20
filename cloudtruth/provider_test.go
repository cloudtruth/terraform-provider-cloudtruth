package cloudtruth

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"cloudtruth": testAccProvider,
	}
}

var (
	testProviderFactories = map[string]func() (*schema.Provider, error){
		"cloudtruth": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_VAR_cloudtruth_api_key"); v == "" {
		t.Fatal("TF_VAR_cloudtruth_api_key must be set for acceptance tests")
	}
}
