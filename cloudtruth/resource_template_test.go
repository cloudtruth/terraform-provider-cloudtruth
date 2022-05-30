package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const templateDesc = "Just a description of a template"
const updateTemplateDesc = "A new description of a template"
const templateValue = "some_var={{ProdOnlyRegularParam}}"
const updateTemplateValue = "some_other_var={{ProdOnlyRegularParam}}"

func TestAccResourceTemplateBasic(t *testing.T) {
	createTemplateName := fmt.Sprintf("TestTemplate-%s", resource.UniqueId())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateCreateBasic(createTemplateName, templateDesc, templateValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "name", createTemplateName),
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "description", templateDesc),
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "value", templateValue),
				),
			},
			{
				Config: testAccResourceTemplateUpdateBasic(createTemplateName, updateTemplateDesc, updateTemplateValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "name", createTemplateName),
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "description", updateTemplateDesc),
					resource.TestCheckResourceAttr("cloudtruth_template.basic", "value", updateTemplateValue),
				),
			},
		},
	})
}

func testAccResourceTemplateCreateBasic(templateName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_template" "basic" {
  		name        = "%s"
  		description = "%s"
		value       = "%s"
	}
	`, templateName, desc, value)
}

func testAccResourceTemplateUpdateBasic(templateName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_template" "basic" {
  		name        = "%s"
  		description = "%s"
        value       = "%s"
	}
	`, templateName, desc, value)
}
