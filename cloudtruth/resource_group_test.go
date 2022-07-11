package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
	"time"
)

const groupDesc = "Just a description of a group"

//const updateGroupDesc = "A new description of a group"

func TestAccResourceGroupBasic(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserCreateBasic(groupName, groupDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "name", groupName),
					resource.TestCheckResourceAttr("cloudtruth_group.basic", "description", groupDesc),
				),
			},
		},
	})
}

func testAccResourceUserCreateBasic(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "basic" {
  		name = "%s"
  		description  = "%s"
	}
	`, name, desc)
}
