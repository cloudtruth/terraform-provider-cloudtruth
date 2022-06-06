package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

const paramDesc = "Just a description of a parameter"
const updateParamDesc = "A new description of an parameter"
const paramVal = "A useful string"
const updateParamVal = "A new useful string"

func TestAccResourceParameterBasic(t *testing.T) {
	createParamName := fmt.Sprintf("TestParam-%s", resource.UniqueId())
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
						strconv.FormatBool(true)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "evaluate",
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
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "evaluate",
						strconv.FormatBool(true)),
				),
			},
			{
				Config: testAccResourceParameterCreateProdOnly(accTestProject, createParamName, paramDesc, paramVal),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "description", paramDesc),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "value", paramVal),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "external",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "evaluate",
						strconv.FormatBool(true)),
					resource.TestCheckResourceAttr("cloudtruth_parameter.prod_only", "environment", "production"),
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
        secret      = true # start off as secret == true, evaluate == false
        evaluate    = false
	}
	`, projName, paramName, desc, value)
}

// todo external test
func testAccResourceParameterUpdateBasic(projName, paramName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "basic" {
        project     = "%s"
  		name        = "%s"
  		description = "%s"
        value       = "%s"
        secret      = false 
        evaluate    = true # secrets cannot be evaluated/dynamic
	}
	`, projName, paramName, desc, value)
}

// This is a case where we create a parameter using the same name but overriding it in the
// production environment i.e. not inheriting from the default environment
func testAccResourceParameterCreateProdOnly(projName, paramName, desc, value string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "prod_only" {
		project     = "%s"
        environment = "production"
  		name        = "%s"
  		description = "%s"
        value       = "%s"
        secret      = false
        evaluate    = true
	}
	`, projName, paramName, desc, value)
}
