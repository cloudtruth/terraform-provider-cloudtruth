package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const templateDesc = "Just a description of a template"
const updateTemplateDesc = "A new description of a template"
const templateValue = "some_var={{ProdOnlyRegularParam}}"
const updateTemplateValue = "some_other_var={{ProdOnlyRegularParam}}"

func TestAccResourceTemplateBasic(t *testing.T) {
	createTemplateName := fmt.Sprintf("TestTemplate-%s", uuid.New().String())
	fullResourceName := "cloudtruth_template.basic"
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateCreateBasic(accTestProject, createTemplateName, templateDesc, templateValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "name", createTemplateName),
					resource.TestCheckResourceAttr(fullResourceName, "description", templateDesc),
					resource.TestCheckResourceAttr(fullResourceName, "value", templateValue),
				),
			},
			{
				ResourceName:      fullResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: paramOrTemplateStateIDFunc(fullResourceName, accTestProject),
			},
			{
				Config: testAccResourceTemplateUpdateBasic(accTestProject, createTemplateName, updateTemplateDesc, updateTemplateValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "name", createTemplateName),
					resource.TestCheckResourceAttr(fullResourceName, "description", updateTemplateDesc),
					resource.TestCheckResourceAttr(fullResourceName, "value", updateTemplateValue),
				),
			},
		},
	})
}

func testAccResourceTemplateCreateBasic(projName, templateName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_template" "basic" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		value       = "%s"
	}
	`, projName, templateName, desc, value)
}

func testAccResourceTemplateUpdateBasic(projName, templateName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_template" "basic" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
        value       = "%s"
	}
	`, projName, templateName, desc, value)
}
