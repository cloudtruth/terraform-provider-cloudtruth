package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
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
		},
	})
}

func TestAccResourceProjectForceDelete(t *testing.T) {
	forceDeleteProjName := fmt.Sprintf("TestDeleteProject-%s", resource.UniqueId())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		// todo: uncomment when https://github.com/hashicorp/terraform-plugin-sdk/pull/976 has been merged
		// and allows us to confirm that the delete indeed fails, for now we set force_
		//ExpectDestroyError: regexp.MustCompile(`.*cannot be deleted from the CloudTruth provider, you must enable 'force_delete' to allow this deletes.*`),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectForceDeleteDisabled(forceDeleteProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "name", forceDeleteProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "description", desc),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "force_delete",
						strconv.FormatBool(false)),
				),
			},
			{ // todo: move this to a new Test(Step) when we can validate that force_delete prevents deletions
				// keeping it here now ensures that temporary resource cleanup will happen
				Config: testAccResourceProjectForceDeleteEnabled(forceDeleteProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "name", forceDeleteProjName),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "description", desc),
					resource.TestCheckResourceAttr("cloudtruth_project.force_delete", "force_delete",
						strconv.FormatBool(true)),
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
