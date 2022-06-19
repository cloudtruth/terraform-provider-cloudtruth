package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"regexp"
	"testing"
	"time"
)

// todo:
// add tag tests when setup/cleanup scripts are in place i.e. for now minimized the amount and complexity
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

const testAccParametersAsOf = `
data "cloudtruth_parameters" "multi_env" {
  project     = "%s"
  environment = "%s"
  as_of       = "%s"
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
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "environment", "default"),
					testAccCheckParametersValueMap("data.cloudtruth_parameters.multi_env", expEnvValues),
				),
			},
		},
	})
}

func TestAccAsOfDataSourceParameters(t *testing.T) {
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				// Use a date when the parameters did not exist
				Config: fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", "2000-01-01T00:00:00"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "environment", "default"),
					testAccCheckParametersValueMap("data.cloudtruth_parameters.multi_env", expEnvValues),
				),
				ExpectError: regexp.MustCompile("404 Not Found"),
			},
			{
				// And a date when the parameters definitely do exist, for simplicity we use time.Now
				Config: fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", now),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.multi_env", "environment", "default"),
					testAccCheckParametersValueMap("data.cloudtruth_parameters.multi_env", expEnvValues),
				),
			},
		},
	})
}

// Confirm that we have the correct map of parameter values for each environment
// We only check for the existence of expected parameters, we do NOT check that unexpected parameters do NOT exist as
// ephemeral parameters will come and go when resource tests run. We can also expect occasional cruft parameters to linger
// if tests and/or cleanup scripts fail
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
