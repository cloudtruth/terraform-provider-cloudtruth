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
  environment = "%s"
}
`

var expEnvValues = map[string]interface{}{
	"parameter_values.DefaultRegularParam": regularParamVal,
	"parameter_values.DefaultSecretParam":  secretParamVal,
}

func TestAccDataSourceParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameters, project, "default"),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", project),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "environment", "default"),
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
		for k, expVal := range expMap {
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
