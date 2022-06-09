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

func TestDataSourceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
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
