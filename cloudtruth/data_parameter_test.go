package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

// For the sake of isolation/bootstrapping
// The tests in the file assume that the following resources already exist in the target
// CloudTruth account, if they are destroyed/altered these tests will fail
// The resource acceptance tests are "full service" tests which instantiate, modify and destroy
// all of the resources which they reference.
const (
	defaultEnv      = "default"
	project         = "AcceptanceTest"
	regularParam    = "DefaultRegularParam"
	regularParamVal = "notreallyasecret"
	secretParam     = "DefaultSecretParam"
	secretParamVal  = "ultratopsecret"
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
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameter, project, defaultEnv,
					regularParam, project, defaultEnv, secretParam),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "project", project),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "name", regularParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret", "project", project),
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
