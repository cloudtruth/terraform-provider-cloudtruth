package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

func TestAccResourceParameterBasic(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateBasic(accTestProject, createParamName, paramDesc, paramVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", paramDesc),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "value", paramVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "dynamic",
						strconv.FormatBool(false)),
				),
			},
			{
				Config: testAccResourceParameterUpdateBasic(accTestProject, createParamName, updateParamDesc, updateParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", updateParamDesc),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "value", updateParamVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "dynamic",
						strconv.FormatBool(false)),
				),
			},
		},
	})
}

func TestAccResourceParameterProdOverride(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateProdOverride(createParamName, paramVal, createParamName, prodParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_override", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_override", "value", prodParamVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_override", "environment", "production"),
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
				Config: testAccResourceParameterCreateExternal(createParamName, paramDesc, githubLocation, githubFilter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.external", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.external", "description", paramDesc),
					resource.TestCheckResourceAttr("cloudtruth_parameter.external", "location", githubLocation),
					resource.TestCheckResourceAttr("cloudtruth_parameter.external", "filter", githubFilter),
				),
			},
		},
	})
}

func testAccResourceParameterCreateBasic(projName, paramName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
        value       = "%s"
	}
	`, projName, paramName, desc, value)
}

func testAccResourceParameterUpdateBasic(projName, paramName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
        project     = "%s"
  		name        = "%s"
  		description = "%s"
        value       = "%s"
	}
	`, projName, paramName, desc, value)
}

// This is a case where we create a parameter using the same name but overriding it in the
// production environment i.e. not inheriting from the default environment
func testAccResourceParameterCreateProdOverride(defaultParam, value, prodParam, prodValue string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
        project     = "AcceptanceTest"
  		name        = "%s"
  		description = "Just a description of a parameter"
        value       = "%s"
	}	

	resource "cloudtruth_parameter" "prod_override" {
		project     = "AcceptanceTest"
        environment = "production"
  		name        = "%s"
  		description = "Just a description of a parameter"
        value       = "%s"
        depends_on = [
        	cloudtruth_parameter.basic
        ]
	}
	`, defaultParam, value, prodParam, prodValue)
}

func testAccResourceParameterCreateExternal(paramName, desc, location, filter string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "external" {
		project          = "AcceptanceTest"
  		name             = "%s"
  		description      = "%s"
        external         = true
        location         = "%s"
        filter           = "%s"
	}
	`, paramName, desc, location, filter)
}
