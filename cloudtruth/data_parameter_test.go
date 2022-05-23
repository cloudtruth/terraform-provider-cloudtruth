package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

/* For now, these tests will rely on the following pre-canned resources:
   project - AcceptanceTest
   environments - default/development/staging/production
   parameters - DefaultRegularParam, DefaultSecretParam
   Eventually, when resources are in place, it may make more sense to
   set up full circuit tests. . .
*/
const (
	defaultEnv      = "default"
	project         = "AcceptanceTest"
	regularParam    = "DefaultRegularParam"
	regularParamVal = "notreallyasecret"
	secretParam     = "DefaultSecretParam"
	secretParamVal  = "ultratopsecret"
)

// todo: add the secret param test validation
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
			{ // Environment and project specified at the provider level
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
const testAccParameterProvEnvProj = `
data "cloudtruth_parameter" "nonsecret_prov_env_proj" {
  name        = "%s"
}

data "cloudtruth_parameter" "secret_prov_env_proj" {
  name        = "%s"
}
`
