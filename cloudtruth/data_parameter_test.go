package cloudtruth

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceParameter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceParameter,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.foo", "env", "production"),
					resource.TestCheckResourceAttr("data.cloudtruth_parameter.foo", "project", "MyFirstProject"),
				),
			},
		},
	})
}

const testAccDataSourceParameter = `
data "cloudtruth_parameter" "foo" {
  env     = "production"
  name    = "MyFirstParameter"
  project = "MyFirstProject"
}
`
