package cloudtruth

/*
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
	"time"
)

func TestAccResourceAccessGrantGroup(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	projName := fmt.Sprintf("TestProject-%d", rand.Intn(100000))
	contribRole := "CONTRIB"
	adminRole := "ADMIN"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGrantGroup(projName, groupName, testUser2.name, contribRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "group", groupName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "project", projName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "role", contribRole),
				),
			},
			{
				Config: testAccResourceGrantGroup(projName, groupName, testUser2.name, adminRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "group", groupName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "project", projName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "role", adminRole),
				),
			},
		},
	})
}

func TestAccResourceAccessGrantUser(t *testing.T) {
	envName := fmt.Sprintf("TestEnv-%d", rand.Intn(100000))
	contribRole := "CONTRIB"
	adminRole := "ADMIN"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGrantUser(envName, testUser2.name, contribRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "user", testUser2.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "role", contribRole),
				),
			},
			{
				Config: testAccResourceGrantUser(envName, testUser2.name, adminRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "user", testUser2.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.access_grant", "role", adminRole),
				),
			},
		},
	})
}

func testAccResourceGrantGroup(project, groupName, userName, role string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "project" {
  		name        = "%s"
        force_delete = true
	}

	resource "cloudtruth_group" "group" {
  		name  = "%s"
		users = ["%s"]
	}

	resource "cloudtruth_access_grant" "access_grant" {
  		group   = cloudtruth_group.group.name
  		project = cloudtruth_project.project.name
		role    = "%s"
	}
	`, project, groupName, userName, role)
}

func testAccResourceGrantUser(environment, userName, role string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "environment" {
  		name  = "%s"
		force_delete = true
	}

	resource "cloudtruth_access_grant" "access_grant" {
  		user        = "%s"
  		environment = cloudtruth_environment.environment.name
		role        = "%s"
	}
	`, environment, userName, role)
}
*/
