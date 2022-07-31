package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceParameterProdOverride(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateProdOverride(createParamName, createParamName, prodParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "value", prodParamVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "environment", "production"),
				),
			},
		},
	})
}

func TestAccResourceParameterExternal(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreventPostDestroyRefresh: true,
		PreCheck:                  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateExternal(createParamName, paramDesc, createParamName, githubLocation, githubFilter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "location", githubLocation),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "filter", githubFilter),
				),
			},
		},
	})
}

// This is a case where we create a parameter using the same name but overriding it in the
// production environment i.e. not inheriting from the default environment
func testAccResourceParameterCreateProdOverride(name, paramName, prodValue string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
        project     = "AcceptanceTest"
  		name        = "%s"
  		description = "Just a description of a parameter"
	}

	resource "cloudtruth_parameter_value" "prod_override" {
        environment    = "production"
  		parameter_name = "%s"
        value          = "%s"
        depends_on = [
        	cloudtruth_parameter.basic
        ]
	}
	`, name, paramName, prodValue)
}

func testAccResourceParameterCreateExternal(name, desc, paramName, location, filter string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "external" {
        project     = "AcceptanceTest"
  		name        = "%s"
  		description = "%s"
	}

	resource "cloudtruth_parameter_value" "external" {
		project          = "AcceptanceTest"
  		parameter_name   = "%s"
        external         = true
        location         = "%s"
        filter           = "%s"
        depends_on = [
        	cloudtruth_parameter.external
        ]
	}
	`, name, desc, paramName, location, filter)
}
