package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccResourceParameterValueCreateUpdate(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	environment := "staging"
	resName := "create_update"
	fullResourceName := fmt.Sprintf("cloudtruth_parameter_value.%s", resName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreate(createParamName, resName, environment, stagingParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "value", stagingParamVal),
					resource.TestCheckResourceAttr(fullResourceName, "environment", environment),
					resource.TestCheckResourceAttr(fullResourceName, "external", "false"),
				),
			},
			{
				Config: testAccResourceParameterCreate(createParamName, resName, environment, stagingParamValUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "value", stagingParamValUpdate),
					resource.TestCheckResourceAttr(fullResourceName, "environment", environment),
					resource.TestCheckResourceAttr(fullResourceName, "external", "false"),
				),
			},
			{
				ResourceName:      fullResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: paramValueStateIDFunc(fullResourceName, fmt.Sprintf("cloudtruth_parameter.%s", resName), accTestProject, environment),
			},
		},
	})
}

func paramValueStateIDFunc(valueResourceName, paramResourceName, projectName, envName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		paramValueID := s.RootModule().Resources[valueResourceName].Primary.ID
		paramID := s.RootModule().Resources[paramResourceName].Primary.ID
		return fmt.Sprintf("%s.%s.%s.%s", projectName, envName, paramID, paramValueID), nil
	}
}

func TestAccResourceParameterValueOverride(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	overrideEnv := "production"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateValueOverride(createParamName, overrideEnv, prodParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "value", prodParamVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.prod_override", "environment", overrideEnv),
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
				Config: testAccResourceParameterValueCreateExternal(createParamName, paramDesc, githubLocation, githubFilter),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "parameter_name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "location", githubLocation),
					resource.TestCheckResourceAttr("cloudtruth_parameter_value.external", "filter", githubFilter),
				),
			},
		},
	})
}

func TestAccResourceParameterValueChangeEnv(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	startEnv := "default"
	updateEnv := "production"
	paramVal := "some string"
	resName := "update_env"
	fullResourceName := fmt.Sprintf("cloudtruth_parameter_value.%s", resName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreate(createParamName, resName, startEnv, paramVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "value", paramVal),
					resource.TestCheckResourceAttr(fullResourceName, "environment", startEnv),
				),
			},
			{
				Config: testAccResourceParameterCreate(createParamName, resName, updateEnv, prodParamVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "value", prodParamVal),
					resource.TestCheckResourceAttr(fullResourceName, "environment", updateEnv),
				),
			},
			{
				Config: testAccResourceParameterCreate(createParamName, resName, updateEnv, paramVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "parameter_name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "value", paramVal),
					resource.TestCheckResourceAttr(fullResourceName, "environment", updateEnv),
				),
			},
		},
	})
}

func testAccResourceParameterCreate(paramName, resName, envName, value string) string {
	return fmt.Sprintf(`
        resource "cloudtruth_parameter" "%s" {
        		project     = "AcceptanceTest"
				name        = "%s"
				description = "Just a description of a parameter"
        }

        resource "cloudtruth_parameter_value" "%s" {
				project        = "AcceptanceTest"
				environment    = "%s"
                parameter_name = cloudtruth_parameter.%s.name
                value          = "%s"
        }
	`, resName, paramName, resName, envName, resName, value)
}

func testAccResourceParameterCreateValueOverride(name, envName, value string) string {
	return fmt.Sprintf(`
        resource "cloudtruth_parameter" "basic" {
                project     = "AcceptanceTest"
                name        = "%s"
                description = "Just a description of a parameter"
        }

        resource "cloudtruth_parameter_value" "prod_override" {
                project        = "AcceptanceTest"
                environment    = "%s"
  				parameter_name = cloudtruth_parameter.basic.name
                value          = "%s"
        }
	`, name, envName, value)
}

func testAccResourceParameterValueCreateExternal(name, desc, location, filter string) string {
	return fmt.Sprintf(`
        resource "cloudtruth_parameter" "external" {
                project     = "AcceptanceTest"
                name        = "%s"
                description = "%s"
        }

        resource "cloudtruth_parameter_value" "external" {
	        	project          = "AcceptanceTest"
                parameter_name   = cloudtruth_parameter.external.name
                external         = true
                location         = "%s"
                filter           = "%s"
        }
	`, name, desc, location, filter)
}
