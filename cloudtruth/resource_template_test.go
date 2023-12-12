package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const templateDesc = "Just a description of a template"
const updateTemplateDesc = "A new description of a template"
const templateValue = "some_var={{DefaultRegularParam}}"
const updateTemplateValue = "some_other_var={{DefaultRegularParam}}"

/*
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
*/
func TestAccResourceTemplateFull(t *testing.T) {
	defaultParamTemplateName := fmt.Sprintf("TestTemplate-%s", uuid.New().String())
	sinkParamName := fmt.Sprintf("SinkParam-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateWithDataSource(accTestProject, defaultParamTemplateName, sinkParamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_template.default_parameters", "name", defaultParamTemplateName),
					//resource.TestCheckResourceAttr(fullResourceName, "value", templateValue),
				),
			},
			{
				Config: testAccResourceTemplateWithDataSource(accTestProject, defaultParamTemplateName, sinkParamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_template.default_parameters", "name", defaultParamTemplateName),
					//resource.TestCheckResourceAttr(fullResourceName, "value", updateTemplateValue),
				),
			},
		},
	})
}

// Create a template that prints all parameters with the word "Default" in their names
// This harness expects at least 3 such parameters. Other tests do as well, therefore it's important to keep or
// provision these fixtures in the target test project
func testAccResourceTemplateWithDataSource(projName, templateName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_template" "default_parameters" {
        name   			= "%s"
        value          	=  "{%%- for param in cloudtruth.parameters -%%}\n{%% if param[0] contains \"Default\" %%}{{ param[0] }}: {{ param[1] }}\n{%% endif -%%}\n{%%- endfor -%%}"
        project       	= "%s"
	}

	data "cloudtruth_template" "default_parameters" {
        name           = "%s"
        project        = "%s"
        environment    = "default"
        depends_on     = [cloudtruth_template.default_parameters]
	}

	resource "cloudtruth_parameter" "sink" {
		project     = "%s"
		name        = "%s"
		depends_on  = [cloudtruth_template.default_parameters]
	}

    resource "cloudtruth_parameter_value" "sink_value" {
		project        = "%s"
		environment    = "default"
		parameter_name = cloudtruth_parameter.sink.name
		value          = data.cloudtruth_template.default_parameters.value
	}
	`, templateName, projName, templateName, projName, projName, paramName, projName)
}
