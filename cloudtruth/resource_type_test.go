package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const typeDesc = "Just a description of a type"
const updateTypeDesc = "A new description of an type"

func TestAccResourceStringTypeWithThreeRules(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "string_three_rules"
	min, max := 0, 1000
	regex := ".*"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateWithThreeRules(resourceName, createTypeName, typeDesc, min, max, regex),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "name", createTypeName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "description", typeDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "type", "string"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "min", fmt.Sprint(min)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "max", fmt.Sprint(max)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "regex", regex),
				),
			},
		},
	})
}

func TestAccResourceTypeBasic(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "basic"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateBasic(resourceName, createTypeName, typeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "name", createTypeName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "description", typeDesc),
				),
			},
			{
				Config: testAccResourceTypeCreateBasic(resourceName, createTypeName, updateTypeDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "name", createTypeName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "description", updateTypeDesc),
				),
			},
		},
	})
}

func TestAccResourceIntTypeWithTwoRules(t *testing.T) {
	createTypeName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "int_two_rules"
	min, max := 0, 1000
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTypeCreateIntWithTwoRules(resourceName, createTypeName, typeDesc, min, max),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "name", createTypeName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "description", typeDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "type", "integer"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "min", fmt.Sprint(min)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_type.%s", resourceName), "max", fmt.Sprint(max)),
				),
			},
		},
	})
}

func testAccResourceTypeCreateBasic(resourceName, typeName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "%s" {
  		name        = "%s"
  		description = "%s"
		type   = "string"
	}
	`, resourceName, typeName, desc)
}

func testAccResourceTypeCreateIntWithTwoRules(resourceName, typeName, desc string, max, min int) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "%s" {
  		name        = "%s"
  		description = "%s"
		type        = "integer"
		min         = "%d"
		max         = "%d"
	}
	`, resourceName, typeName, desc, max, min)
}

func testAccResourceTypeCreateWithThreeRules(resourceName, typeName, desc string, max, min int, regex string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_type" "%s" {
  		name        = "%s"
  		description = "%s"
		type        = "string"
		min         = "%d"
        max         = "%d"
    	regex       = "%s"
	}
	`, resourceName, typeName, desc, max, min, regex)
}
