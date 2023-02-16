package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
	"time"
)

const testAccParameter = `
data "cloudtruth_parameter_value" "nonsecret" {
  project        = "%s"
  environment    = "%s"
  parameter_name = "%s"
}
data "cloudtruth_parameter_value" "secret" {
  project        = "%s"
  environment    = "%s"
  parameter_name = "%s"
}
`

// project and environment must be specified via
// the CLOUDTRUTH_PROJECT and CLOUDTRUTH_ENVIRONMENT environment variables
// which will set them at the provider level
const testAccParameterProvEnvProj = `
data "cloudtruth_parameter_value" "nonsecret_param" {
  parameter_name = "%s"
}

data "cloudtruth_parameter_value" "secret_param" {
  parameter_name = "%s"
}

data "cloudtruth_parameter_value" "external_param" {
  parameter_name = "%s"
}
`

const testAccMissingParameter = `
data "cloudtruth_parameter_value" "missing" {
  parameter_name = "idontexist"
}
`

const testAccParameters = `
data "cloudtruth_parameter_values" "multi_env" {
  project     = "%s"
  environment = "%s"
}
`

const testAccParametersAsOf = `
data "cloudtruth_parameter_values" "as_of_params" {
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
				ExpectError: regexp.MustCompile("could not find the parameter"),
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
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.nonsecret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.nonsecret", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.nonsecret", "parameter_name", regularParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.nonsecret", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.secret", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.secret", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.secret", "parameter_name", secretParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.secret", "value", secretParamVal),
				),
			},
			{ // In this step, the environment and project default to the values specified at the provider level
				Config: fmt.Sprintf(testAccParameterProvEnvProj, regularParam, secretParam, regularExternalParam),
				Check: resource.ComposeTestCheckFunc(
					// regular parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.nonsecret_param", "value", regularParamVal),

					// secret parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.secret_param", "value", secretParamVal),

					// regular external parameter
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_value.external_param", "value", regularExternalParamVal),
				),
			},
		},
	})
}

func TestAccDataSourceParameterValues(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccParameters, accTestProject, "default"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_values.multi_env", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_values.multi_env", "environment", "default"),
					testAccCheckDataSourceValueMap("data.cloudtruth_parameter_values.multi_env", expParameterValues),
				),
			},
		},
	})
}

func TestAccAsOfDataSourceParameterValues(t *testing.T) {
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				// Use a date when the parameters did not exist
				Config:      fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", "2000-01-01T00:00:00"),
				ExpectError: regexp.MustCompile("404 client error.*No ProjectLedger matches the given query"),
			},
			{
				// And a date when the parameters definitely do exist, for simplicity we use time.Now
				Config: fmt.Sprintf(testAccParametersAsOf, accTestProject, "default", now),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_values.as_of_params", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter_values.as_of_params", "environment", "default"),
					testAccCheckDataSourceValueMap("data.cloudtruth_parameter_values.as_of_params", expParameterValues),
				),
			},
		},
	})
}
