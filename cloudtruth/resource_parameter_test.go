package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const paramDesc = "Just a description of a parameter"
const updateParamDesc = "A new description of an parametr"

// todo: add a non-default parent test
// and re-parent test if applicable
func TestAccResourceParameterBasic(t *testing.T) {
	createParamName := fmt.Sprintf("TestParam-%s", resource.UniqueId())
	updateParamName := fmt.Sprintf("updated-%s", createParamName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateBasic(createParamName, paramDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", paramDesc),
				),
			},
			{
				Config: testAccResourceParameterUpdateBasic(updateParamName, updateEnvDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", updateParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", updateParamDesc),
				),
			},
		},
	})
}

func testAccResourceParameterCreateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, envName, desc)
}

func testAccResourceParameterUpdateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, envName, desc)
}
