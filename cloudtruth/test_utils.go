package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// For the sake of isolation/bootstrapping
// The provider acceptance tests assume that the following resources already exist in the target
// CloudTruth account, if they are destroyed/altered these tests will fail
// The resource acceptance tests are "full service" tests which instantiate, modify and destroy
// the resources which they reference.

// The following constants are shared across multiple tests
// Constants and variables which are unique per resource/data source type are located in the
// corresponding *_test.go files
const (
	accTestAWSIntegrationName   = "my_cloudtruth_role@426203803362"
	accTestAzureIntegrationName = "cttesting - cfafeccc-5946-4cb3-a2bc-b6a3961c2ffa"
	defaultEnv                  = "default"
	accTestProject              = "AcceptanceTest"
	regularParam                = "DefaultRegularParam"
	regularParamVal             = "notreallyasecret"
	secretParam                 = "DefaultSecretParam"
	secretParamVal              = "ultratopsecret"
	regularExternalParam        = "DefaultRegularExternalParam"
	regularExternalParamVal     = "5.0"
	paramDesc                   = "Just a description of a parameter"
	updateParamDesc             = "A new description of an parameter"
	genericDesc                 = "A generic description"
	stagingParamVal             = "A useful string only in staging"
	stagingParamValUpdate       = "An updated & useful string only in staging"
	prodParamVal                = "A useful string only in production"
	basicTemplate               = "BasicTemplate"
	basicTemplateVal            = "Regular Parameter: notreallyasecret, Secret Parameter ultratopsecret"
	defaultTemplateVal          = "notreallyasecret"
	githubLocation              = "github://cloudtruth/terraform-provider-cloudtruth/main/terraform-registry-manifest.json"
	githubFilter                = "metadata.protocol_versions[0]"
	epochTimeTag                = "EpochTime"
	epochTimeTagTimestamp       = "1970-01-01T05:00:00Z"
	ciServiceAccountName        = "ACCEPTANCE_TEST_TOKEN" // used in tests and CI
	ciServiceAccountRole        = "ADMIN"
)

// Utility function to use with data sources and resources that are maps
// Note that this function only checks the existence of expected key/value pairs
// The acceptance tests create ephemeral resources which means that occasionally
// lingering resources may exist in the CloudTruth Acceptance Test account
// when tests and/or cleanup scripts fail
func testAccCheckDataSourceValueMap(resourceName string, expMap map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		for k, expVal := range expMap {
			actVal, ok := rs.Primary.Attributes[k]
			if !ok {
				return fmt.Errorf("did not find expected value: %s", k)
			}
			if expVal != actVal {
				return fmt.Errorf("expected value: %s, actual value %s", expVal, actVal)
			}
		}
		return nil
	}
}
