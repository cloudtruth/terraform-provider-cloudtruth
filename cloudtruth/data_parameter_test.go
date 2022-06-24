package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
data "cloudtruth_parameter" "nonsecret_param" {
  name        = "%s"
}

data "cloudtruth_parameter" "secret_param" {
  name        = "%s"
}
`

const testAccMissingParameter = `
data "cloudtruth_parameter" "missing" {
  name     = "idontexist"
}
`

const testAccParameters = `
data "cloudtruth_parameters" "multi_env" {
  project     = "%s"
  environment = "%s"
}
`

const testAccParametersAsOf = `
data "cloudtruth_parameters" "as_of_params" {
  project     = "%s"
  environment = "%s"
  as_of       = "%s"
}
`

var expParameterValues = map[string]interface{}{
	"parameter_values.DefaultRegularParam": regularParamVal,
	"parameter_values.DefaultSecretParam":  secretParamVal,
}

func TestMissingDataSourceParameter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config:      testAccMissingParameter,
				ExpectError: regexp.MustCompile("expected 1 value for parameter"),
			},
		},
	})
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
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.nonsecret_param", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.secret_param", "value", secretParamVal),
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
					testAccCheckDataSourceValueMap("data.cloudtruth_parameters.multi_env", expParameterValues),
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
				Config:      fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", "2000-01-01T00:00:00"),
				ExpectError: regexp.MustCompile("404 Not Found"),
			},
			{
				// And a date when the parameters definitely do exist, for simplicity we use time.Now
				Config: fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", now),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.as_of_params", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameters.as_of_params", "environment", "default"),
					testAccCheckDataSourceValueMap("data.cloudtruth_parameters.as_of_params", expParameterValues),
				),
			},
		},
	})
}
