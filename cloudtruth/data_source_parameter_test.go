package cloudtruth

import (
//"regexp"
//"testing"

//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/*func TestDataSourceParameter(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceParameter,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.parameter_data_source.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}
*/
const testAccDataSourceParameter = `
data "parameter_data_source" "foo" {
  env     = "production"
  project = "MyFirstProject"
  wrap    = true
}
`
