package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
	"time"
)

const groupDesc = "Just a description of a group"
const updateGroupDesc = "A new description of a group"

// Use internal accounts since they guaranteed to exist
// Need to document/support a way for external contributors to
// specify their own users
const groupTestUser1 = "matthewcummings516@gmail.com"
const groupTestUser2 = "matt@cloudtruth.com"

func TestAccResourceGroupBasic(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateBasic(groupName, groupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "description", groupDesc),
				),
			},
			{
				Config: testAccResourceGroupUpdateBasic(groupName, updateGroupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "description", updateGroupDesc),
				),
			},
		},
	})
}

func TestAccResourceGroupWithUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateWithUser(groupName, groupDesc, groupTestUser1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "description", groupDesc),
					//resource.T("cloudtruth_group.user_test", "users", []string{groupTestUser1}),
				),
			},
			{ // Change group membership & description
				Config: testAccResourceGroupCreateWithoutUser(groupName, updateGroupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "description", updateGroupDesc),
					//resource.TestCheckR("cloudtruth_group.user_test", "users", []string{groupTestUser1}),
				),
			},
		},
	})
}

func TestAccResourceGroupWithUsers(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateWithUsers(groupName, groupDesc, groupTestUser1, groupTestUser2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "description", groupDesc),
					//resource.T("cloudtruth_group.multi_user", "users", []string{groupTestUser1}),
				),
			},
			{ // Reverse the order of names in the list
				Config: testAccResourceGroupCreateWithUsers(groupName, groupDesc, groupTestUser2, groupTestUser1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "name", groupName),
					//resource.TestCheckR("cloudtruth_group.multi_user", "users", []string{groupTestUser1}),
				),
			},
		},
	})
}

func testAccResourceGroupCreateBasic(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "basic" {
  		name = "%s"
  		description  = "%s"
	}
	`, name, desc)
}

func testAccResourceGroupUpdateBasic(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "basic" {
  		name = "%s"
  		description  = "%s"
	}
	`, name, desc)
}

func testAccResourceGroupCreateWithUser(name, desc, user string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "user_test" {
  		name        = "%s"
  		description = "%s"
        users       = ["%s"]
	}
	`, name, desc, user)
}

func testAccResourceGroupCreateWithoutUser(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "user_test" {
  		name        = "%s"
  		description = "%s"
        users       = []
	}
	`, name, desc)
}

func testAccResourceGroupCreateWithUsers(name, desc, user1, user2 string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "multi_user" {
  		name        = "%s"
  		description = "%s"
        users       = ["%s", "%s"]
	}
	`, name, desc, user1, user2)
}
