package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const desc = "Just a description of a project"
const updateDesc = "A new description of a project"

func TestAccResourceProjectBasic(t *testing.T) {
	createProjName := fmt.Sprintf("TestProject-%s", resource.UniqueId())
	updateProjName := fmt.Sprintf("updated-%s", createProjName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectCreateBasic(createProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "name", createProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "description", desc),
				),
			},
			{
				Config: testAccResourceProjectUpdateBasic(updateProjName, updateDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "name", updateProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "description", updateDesc),
				),
			},
		},
	})
}

func testAccResourceProjectCreateBasic(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, projName, desc)
}

func testAccResourceProjectUpdateBasic(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, projName, desc)
}
