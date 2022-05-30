package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const paramDesc = "Just a description of a parameter"
const updateParamDesc = "A new description of an parameter"

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
				Config: testAccResourceParameterCreateBasic(accTestProject, createParamName, paramDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", paramDesc),
				),
			},
			{
				Config: testAccResourceParameterUpdateBasic(accTestProject, updateParamName, updateEnvDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", updateParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", updateParamDesc),
				),
			},
		},
	})
}

func testAccResourceParameterCreateBasic(projName, paramName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
	}
	`, projName, paramName, desc)
}

func testAccResourceParameterUpdateBasic(projName, paramName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
        project     = "%s"
  		name        = "%s"
  		description = "%s"
	}
	`, projName, paramName, desc)
}
