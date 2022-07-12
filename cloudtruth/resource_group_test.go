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

// For convenience, reuse the API token used in CI for the
// acceptance tests
// const tokenUser = "ACCEPTANCE_TEST_TOKEN"

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

/*
func TestAccResourceGroupWithUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupCreateWithUser(groupName, groupDesc, tokenUser),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "description", groupDesc),
					//resource.TestCheckR("cloudtruth_group.basic", "users", []string{tokenUser}),
				),
			},
			{
				Config: testAccResourceGroupCreateWithoutUser(groupName, updateGroupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "description", updateGroupDesc),
					//resource.TestCheckR("cloudtruth_group.basic", "users", []string{tokenUser}),
				),
			},
		},
	})
}
*/

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

/*
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
}*/
