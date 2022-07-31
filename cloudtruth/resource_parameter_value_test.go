package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

//func testAccResourceParameterCreate(name, paramName, resName, env, value string) string {

func TestAccResourceParameterValueCreateUpdate(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	environment := "staging"
	resName := "create_update"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreate(createParamName, resName, environment, stagingParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "value", stagingParamVal),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "environment", environment), resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "environment", environment),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "external", "false"),
				),
			},
			{
				Config: testAccResourceParameterCreate(createParamName, resName, environment, stagingParamValUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "value", stagingParamValUpdate),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "environment", environment), resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "environment", environment),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter_value.%s", resName), "external", "false"),
				),
			},
		},
	})
}

func TestAccResourceParameterValueOverride(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateValueOverride(createParamName, createParamName, prodParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "value", prodParamVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "environment", "production"),
				),
			},
		},
	})
}

func TestAccResourceParameterValueExternal(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreventPostDestroyRefresh: true,
		PreCheck:                  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterValueCreateExternal(createParamName, paramDesc, createParamName, githubLocation, githubFilter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "location", githubLocation),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "filter", githubFilter),
				),
			},
		},
	})
}

func testAccResourceParameterCreate(paramName, resName, env, value string) string {
	return fmt.Sprintf(`
        resource "cloudtruth_parameter" "%s" {
        project     = "AcceptanceTest"
                name        = "%s"
                description = "Just a description of a parameter"
        }

        resource "cloudtruth_parameter_value" "%s" {
                project        = "AcceptanceTest"
                environment    = "%s"
                parameter_name = "%s"
                value          = "%s"
                depends_on = [
                      cloudtruth_parameter.%s
                ]
        }
	`, resName, paramName, resName, env, paramName, value, resName)
}

func testAccResourceParameterCreateValueOverride(name, paramName, prodValue string) string {
	return fmt.Sprintf(`
        resource "cloudtruth_parameter" "basic" {
                project     = "AcceptanceTest"
                name        = "%s"
                description = "Just a description of a parameter"
        }

        resource "cloudtruth_parameter_value" "prod_override" {
                project        = "AcceptanceTest"
                environment    = "production"
  		parameter_name = "%s"
                value          = "%s"
                depends_on = [
                      cloudtruth_parameter.basic
                ]
        }
	`, name, paramName, prodValue)
}

func testAccResourceParameterValueCreateExternal(name, desc, paramName, location, filter string) string {
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
