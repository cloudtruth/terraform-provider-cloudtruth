package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const testAccTemplate = `
data "cloudtruth_template" "basic" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}
`

const testAccTemplates = `
data "cloudtruth_templates" "basic_default" {
  project     = "%s"
  environment = "%s"
}
`

var expTemplateValues = map[string]interface{}{
	"template_values.BasicTemplate":   basicTemplateVal,
	"template_values.DefaultTemplate": defaultTemplateVal,
}

func TestDataSourceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccTemplate, accTestProject, defaultEnv, basicTemplate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_template.basic", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_template.basic", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_template.basic", "name", basicTemplate),
					resource.TestCheckResourceAttr("data.cloudtruth_template.basic", "value", basicTemplateVal),
				),
			},
		},
	})
}

func TestDataSourceTemplates(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccTemplates, accTestProject, defaultEnv),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_templates.basic_default", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_templates.basic_default", "project", accTestProject),
					testAccCheckDataSourceValueMap("data.cloudtruth_templates.basic_default", expTemplateValues),
				),
			},
		},
	})
}
