package cloudtruth

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
				Config: fmt.Sprintf(testAccDataSourceParameter, project, defaultEnv,
					regularParam, project, defaultEnv, secretParam),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.regular_param", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.regular_param", "project", project),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.regular_param", "name", regularParam),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.regular_param", "value", regularParamVal),
				),
			},
		},
	})
}

const testAccDataSourceParameter = `
data "cloudtruth_parameter" "regular_param" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}

data "cloudtruth_parameter" "secret_param" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}

`
