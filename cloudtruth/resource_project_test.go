package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const desc = "Just a description"

func TestAccResourceProjectBasic(t *testing.T) {
	basicProjName := fmt.Sprintf("TestProject-%s", resource.UniqueId())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectConfigBasic(basicProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "name", basicProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "description", desc),
				),
			},
		},
	})
}

func testAccResourceProjectConfigBasic(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, projName, desc)
}
