package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
	"testing"
)

const desc = "Just a description of a project"
const updateDesc = "A new description of a project"

// todo: add a non-default parent test
// and re-parent test if applicable

func testAccCheckFoo(s *terraform.State) error {
	fmt.Println("here we are")
	return nil
}

func TestAccResourceProjectBasic(t *testing.T) {
	createProjName := fmt.Sprintf("TestProject-%s", resource.UniqueId())
	updateProjName := fmt.Sprintf("updated-%s", createProjName)
	forceDeleteProjName := fmt.Sprintf("TestDeleteProject-%s", resource.UniqueId())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		CheckDestroy:      testAccCheckFoo,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectCreateBasic(createProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "name", createProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "description", desc),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "force_delete",
						strconv.FormatBool(true)),
				),
			},
			{
				Config: testAccResourceProjectUpdateBasic(updateProjName, updateDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "name", updateProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "description", updateDesc),
					resource.TestCheckResourceAttr("cloudtruth_project.basic", "force_delete",
						strconv.FormatBool(true)),
				),
			},
			{ // The delete should fail
				Config: testAccResourceProjectForceDeleteDisabled(forceDeleteProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "name", forceDeleteProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "description", desc),
				),
				//ExpectError: regexp.MustCompile(`.*cannot be deleted from the CloudTruth provider, you must enable 'force_delete' to allow this deletes.*`),
			},
			{ // Then set force_delete = true and there should be no errors
				Config: testAccResourceProjectForceDeleteEnabled(forceDeleteProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "name", forceDeleteProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "description", desc),
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
        force_delete = true
	}
	`, projName, desc)
}

func testAccResourceProjectUpdateBasic(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "basic" {
  		name         = "%s"
  		description  =  "%s"
		force_delete = true
	}
	`, projName, desc)
}

func testAccResourceProjectForceDeleteDisabled(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "force_delete" {
  		name         = "%s"
  		description  =  "%s"
	}
	`, projName, desc)
}

func testAccResourceProjectForceDeleteEnabled(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "force_delete" {
  		name         = "%s"
  		description  =  "%s"
		force_delete = true
	}
	`, projName, desc)
}
