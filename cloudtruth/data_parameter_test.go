package cloudtruth

import (
	"regexp"
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
					resource.TestMatchResourceAttr("data.cloudtruth_parameter.foo", "env", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceParameter = `
data "cloudtruth_parameter" "foo" {
  env     = "production"
  project = "MyFirstProject"
  wrap    = true
}
`
