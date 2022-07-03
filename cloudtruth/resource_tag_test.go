package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const tagDesc = "Just a description of a tag"
const updateTagDesc = "A new description of an tag"

func TestAccResourceTagBasic(t *testing.T) {
	createTagName := fmt.Sprintf("Test-%s", uuid.New().String())
	timestamp := "2022-07-03T22:01:24Z"
	updateTimestamp := "2022-07-03T22:02:24Z"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTagCreateBasic(createTagName, tagDesc, timestamp),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "name", createTagName),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "description", tagDesc),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "timestamp", timestamp),
				),
			},
			{
				Config: testAccResourceTagUpdateBasic(createTagName, updateTagDesc, updateTimestamp),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "name", createTagName),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "description", updateTagDesc),
					resource.TestCheckResourceAttr("cloudtruth_tag.basic", "timestamp", updateTimestamp),
				),
			},
		},
	})
}

func testAccResourceTagCreateBasic(tagName, desc, timestamp string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_tag" "basic" {
  		name        = "%s"
  		description = "%s"
        timestamp       = "%s"
	}
	`, tagName, desc, timestamp)
}

func testAccResourceTagUpdateBasic(tagName, desc, timestamp string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_tag" "basic" {
  		name        = "%s"
  		description = "%s"
        timestamp       = "%s"
	}
	`, tagName, desc, timestamp)
}
