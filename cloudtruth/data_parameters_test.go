package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccParameters = `
data "cloudtruth_parameters" "multi_env" {
  project     = "%s"
  name        = "%s"
}
`

var expEnvValues = map[string]interface{}{
	"environment_values.default":     regularParamVal,
	"environment_values.development": regularParamVal,
	"environment_values.staging":     regularParamVal,
	"environment_values.production":  regularParamVal,
}

func TestDataSourceParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameters, project, regularParam),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", project),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "name", regularParam),
					testAccCheckParametersValueMap("data.cloudtruth_parameters.multi_env", expEnvValues),
				),
			},
		},
	})
}

// Confirm that we have the correct map of parameter values for each environment
func testAccCheckParametersValueMap(resourceName string, expMap map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		for k, expVal := range expEnvValues {
			actVal, ok := rs.Primary.Attributes[k]
			if !ok {
				return fmt.Errorf("did not find expected parameter value: %s", k)
			}
			if expVal != actVal {
				return fmt.Errorf("expected parameter value: %s, actual parameter value %s", expVal, actVal)
			}
		}
		return nil
	}
}
