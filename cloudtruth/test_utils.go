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
	defaultEnv         = "default"
	accTestProject     = "AcceptanceTest"
	regularParam       = "DefaultRegularParam"
	regularParamVal    = "notreallyasecret"
	secretParam        = "DefaultSecretParam"
	secretParamVal     = "ultratopsecret"
	basicTemplate      = "BasicTemplate"
	basicTemplateVal   = "Regular Parameter: notreallyasecret, Secret Parameter ultratopsecret"
	defaultTemplate    = "DefaultTemplate"
	defaultTemplateVal = "notreallyasecret"
)

// Utility function to use with data sources and resources that are maps
// Note that this function only checks the existence of expected key/value pairs
// The acceptance tests create ephemeral resources which means that occasionally
// lingering cruft resources may exist in the CloudTruth Acceptance Test account
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
