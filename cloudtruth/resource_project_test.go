package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const desc = "Just a description of a project"
const updateDesc = "A new description of a project"

func TestAccResourceProjectBasic(t *testing.T) {
	createProjName := fmt.Sprintf("TestProject-%s", uuid.New().String())
	updateProjName := fmt.Sprintf("updated-%s", createProjName)
	resourceName := "cloudtruth_project.basic"
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProjectCreateBasic(createProjName, desc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", createProjName),
					resource.TestCheckResourceAttr(resourceName, "description", desc),
				),
			},
			/* This is leading to occasional timing issues which appear to be only affecting the tests
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},*/
			{
				Config: testAccResourceProjectUpdateBasic(updateProjName, updateDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateProjName),
					resource.TestCheckResourceAttr(resourceName, "description", updateDesc),
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
	}
	`, projName, desc)
}

func testAccResourceProjectUpdateBasic(projName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "basic" {
  		name         = "%s"
  		description  =  "%s"
	}
	`, projName, desc)
}

func testAccResourceProjectNested(projectOne, projectTwo, projectThree string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "project_one" {
  		name         = "%s"
	}

	resource "cloudtruth_project" "project_two" {
  		name         = "%s"
		parent       = cloudtruth_project.project_one.name
	}

	resource "cloudtruth_project" "project_three" {
  		name         = "%s"
		parent       = cloudtruth_project.project_two.name
	}
	`, projectOne, projectTwo, projectThree)
}
