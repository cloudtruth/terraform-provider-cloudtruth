package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const typeDesc = "Just a description of a type"
const updateTypeDesc = "A new description of an type"

func TestAccResourceTypeBasic(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateBasic(createTypeName, typeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_type.basic", "name", createTypeName),
					resource.TestCheckResourceAttr("cloudtruth_type.basic", "description", typeDesc),
				),
			},
			{
				Config: testAccResourceTypeUpdateBasic(createTypeName, updateTypeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_type.basic", "name", createTypeName),
					resource.TestCheckResourceAttr("cloudtruth_type.basic", "description", updateTypeDesc),
				),
			},
		},
	})
}

func TestAccResourceTypeWithRule(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateWithRule(createTypeName, typeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_type.with_rule", "name", createTypeName),
					resource.TestCheckResourceAttr("cloudtruth_type.with_rule", "description", typeDesc),
				),
			},
		},
	})
}

func TestAccResourceTypeWithThreeRules(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateWithThreeRules(createTypeName, typeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_type.with_three_rules", "name", createTypeName),
					resource.TestCheckResourceAttr("cloudtruth_type.with_three_rules", "description", typeDesc),
				),
			},
		},
	})
}

func testAccResourceTypeCreateBasic(typeName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "basic" {
  		name        = "%s"
  		description = "%s"
		base_type   = "string"
	}
	`, typeName, desc)
}

func testAccResourceTypeUpdateBasic(tagName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "basic" {
  		name        = "%s"
  		description = "%s"
		base_type   = "string"
	}
	`, tagName, desc)
}

func testAccResourceTypeCreateWithRule(typeName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "with_rule" {
  		name        = "%s"
  		description = "%s"
		base_type   = "string"
		rule {
			type       = "min_len"
			constraint = "10"
		}
	}
	`, typeName, desc)
}

func testAccResourceTypeCreateWithThreeRules(typeName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "with_three_rules" {
  		name        = "%s"
  		description = "%s"
		base_type   = "string"
		rule {
			type       = "min_len"
			constraint = "1"
		}
		rule {
			type       = "max_len"
			constraint = "10"
		}
		rule {
			type       = "regex"
			constraint = ".*"
		}
	}
	`, typeName, desc)
}
