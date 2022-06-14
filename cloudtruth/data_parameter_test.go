package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccParameter = `
data "cloudtruth_parameter" "nonsecret" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}
data "cloudtruth_parameter" "secret" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}
`

// project and environment must be specified via
// the CLOUDTRUTH_PROJECT and CLOUDTRUTH_ENVIRONMENT environment variables
// which will set them at the provider level
const testAccParameterProvEnvProj = `
data "cloudtruth_parameter" "nonsecret_prov_env_proj" {
  name        = "%s"
}

data "cloudtruth_parameter" "secret_prov_env_proj" {
  name        = "%s"
}
`

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

func TestDataSourceParameter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameter, accTestProject, defaultEnv,
					regularParam, accTestProject, defaultEnv, secretParam),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "name", regularParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "name", secretParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "value", secretParamVal),
				),
			},
			{ // In this step, the environment and project default to the values specified at the provider level
				Config: fmt.Sprintf(testAccParameterProvEnvProj, regularParam, secretParam),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret_prov_env_proj", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret_prov_env_proj", "value", secretParamVal),
				),
			},
		},
	})
}

func TestAccDataSourceParameters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameters, accTestProject, "default"),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "environment", "default"),
					testAccCheckParametersValueMap("data.cloudtruth_parameters.multi_env", expEnvValues),
				),
			},
		},
	})
}

// Confirm that we have the correct map of parameter values for each environment
// We only check for these parameters, other ephemeral parameters will exist when resource tests run
func testAccCheckParametersValueMap(resourceName string, expMap map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
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
