package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
)

func TestAccResourceAccessGrantGroup(t *testing.T) {
	groupName := fmt.Sprintf("TestGroup-%s", uuid.New().String())
	projName := fmt.Sprintf("TestProject-%s", uuid.New().String())
	contribRole := "CONTRIB"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGrantGroup(projName, groupName, testUser1.name, testUser2.name, contribRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.group_access_grant", "group", groupName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.group_access_grant", "project", projName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.group_access_grant", "role", contribRole),
				),
				SkipFunc: isRunningSelfHosted,
			},
		},
	})
}

func TestAccResourceAccessGrantUser(t *testing.T) {
	envName := fmt.Sprintf("TestEnv-%d", rand.Intn(100000))
	contribRole := "CONTRIB"
	ownerRole := "OWNER"
	viewerRole := "VIEWER"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGrantUser(envName, testUser2.name, testUser1.name, contribRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "user", testUser2.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "role", ownerRole),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "user", testUser1.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "role", contribRole),
				),
				SkipFunc: isRunningSelfHosted,
			},
			{
				Config: testAccResourceGrantUser(envName, testUser2.name, testUser1.name, viewerRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "user", testUser2.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.owner_access_grant", "role", ownerRole),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "user", testUser1.name),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "environment", envName),
					resource.TestCheckResourceAttr("cloudtruth_access_grant.non_owner_access_grant", "role", viewerRole),
				),
				SkipFunc: isRunningSelfHosted,
			},
		},
	})
}

func testAccResourceGrantGroup(project, groupName, groupUser, owner, role string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_project" "project" {
  		name        = "%s"
	}

	resource "cloudtruth_group" "group" {
  		name  = "%s"
		users = ["%s"]
	}

	resource "cloudtruth_access_grant" "owner_access_grant" {
  		user        = "%s"
		project     = cloudtruth_project.project.name
		role        = "OWNER"
	}

	resource "cloudtruth_access_grant" "group_access_grant" {
  		group   = cloudtruth_group.group.name
  		project = cloudtruth_project.project.name
		role    = "%s"
		depends_on = [
			cloudtruth_access_grant.owner_access_grant
		]
	}
	`, project, groupName, groupUser, owner, role)
}

func testAccResourceGrantUser(environment, owner, nonOwner, role string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "environment" {
  		name  = "%s"
	}

	resource "cloudtruth_access_grant" "owner_access_grant" {
  		user        = "%s"
  		environment = cloudtruth_environment.environment.name
		role        = "OWNER"
	}

	resource "cloudtruth_access_grant" "non_owner_access_grant" {
  		user        = "%s"
  		environment = cloudtruth_environment.environment.name
		role        = "%s"
		depends_on = [
			cloudtruth_access_grant.owner_access_grant
		]
	}
	`, environment, owner, nonOwner, role)
}
