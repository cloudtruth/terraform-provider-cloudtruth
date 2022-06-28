package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

const desc = "Just a description of a project"
const updateDesc = "A new description of a project"

func TestAccResourceProjectBasic(t *testing.T) {
	createProjName := fmt.Sprintf("TestProject-%s", uuid.New().String())
	updateProjName := fmt.Sprintf("updated-%s", createProjName)
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
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
	forceDeleteProjName := fmt.Sprintf("TestDeleteProject-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		// todo: uncomment when https://github.com/hashicorp/terraform-plugin-sdk/pull/976 has been merged
		// and allows us to confirm that the delete indeed fails, for now we set force_delet to true
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

func TestAccResourceProjectNested(t *testing.T) {
	projectOneName := fmt.Sprintf("TestProject1-%s", uuid.New().String())
	projectTwoName := fmt.Sprintf("TestProject2-%s", uuid.New().String())
	projectThreeName := fmt.Sprintf("TestProject3-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectNested(projectOneName, projectTwoName, projectThreeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_project.project_one", "name", projectOneName),
					resource.TestCheckResourceAttr("cloudtruth_project.project_two", "name", projectTwoName),
					resource.TestCheckResourceAttr("cloudtruth_project.project_three", "name", projectThreeName),
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

func testAccResourceProjectNested(projectOne, projectTwo, projectThree string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "project_one" {
  		name         = "%s"
		force_delete = true
	}

	resource "cloudtruth_project" "project_two" {
  		name         = "%s"
		parent       = "cloudtruth_project.project_one"
		force_delete = true
	}

	resource "cloudtruth_project" "project_three" {
  		name         = "%s"
		parent       = "cloudtruth_project.project_three"
		force_delete = true
	}
	`, projectOne, projectTwo, projectThree)
}
