package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const groupDesc = "Just a description of a group"
const updateGroupDesc = "A new description of a group"

// Use internal accounts since they guaranteed to exist
// Need to document/support a way for external contributors to
// specify their own users
const groupTestUser1 = "matt@cloudtruth.com"
const groupTestUser2 = "matthewcummings516@gmail.com"

func TestAccResourceGroupBasic(t *testing.T) {
	groupName := fmt.Sprintf("TestGroup-%d", randInt(100000))
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
	groupName := fmt.Sprintf("TestGroup-%d", randInt(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateWithUser(groupName, groupDesc, groupTestUser1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "description", groupDesc),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "users.0", groupTestUser1),
				),
			},
			{ // swap users
				Config: testAccResourceGroupCreateWithUser(groupName, groupDesc, groupTestUser2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "description", groupDesc),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "users.0", groupTestUser2),
				),
			},
			{ // Remove the user & update the description
				Config: testAccResourceGroupCreateWithoutUser(groupName, updateGroupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "description", updateGroupDesc),
					resource.TestCheckResourceAttr("cloudtruth_group.user_test", "users.#", "0"),
				),
			},
		},
	})
}

func TestAccResourceGroupWithUsers(t *testing.T) {
	groupName := fmt.Sprintf("TestGroup-%d", randInt(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateWithUsers(groupName, groupDesc, groupTestUser1, groupTestUser2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "description", groupDesc),
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "users.#", "2"),
					resource.TestCheckResourceAttrSet("cloudtruth_group.multi_user", "users.0"),
					resource.TestCheckResourceAttrSet("cloudtruth_group.multi_user", "users.1"),
				),
			},
			{ // Reverse the order of names in the list, test should still pass
				Config: testAccResourceGroupCreateWithUsers(groupName, groupDesc, groupTestUser2, groupTestUser1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "description", groupDesc),
					resource.TestCheckResourceAttr("cloudtruth_group.multi_user", "users.#", "2"),
					resource.TestCheckResourceAttrSet("cloudtruth_group.multi_user", "users.0"),
					resource.TestCheckResourceAttrSet("cloudtruth_group.multi_user", "users.1"),
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
