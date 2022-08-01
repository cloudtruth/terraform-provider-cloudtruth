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
	resourceName := "basic"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateBasic(accTestProject, resourceName, createParamName, paramDesc, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "name", createParamName),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "description", paramDesc),
					resource.TestCheckResourceAttr("cloudtruth_parameter.basic", "secret",
						strconv.FormatBool(true)),
				),
			},
			{ // Update test, changing the description and the secret toggle
				Config: testAccResourceParameterCreateBasic(accTestProject, resourceName, createParamName, updateParamDesc, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "description", updateParamDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "secret",
						strconv.FormatBool(false)),
				),
			},
		},
	})
}

func TestAccResourceParameterWithRules(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "with_rules"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateWithRules(accTestProject, resourceName, createParamName, paramDesc, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "secret",
						strconv.FormatBool(false)),
				),
			},
		},
	})
}

func testAccResourceParameterCreateBasic(projName, resourceName, paramName, desc string, isSecret bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
	}
	`, resourceName, projName, paramName, desc, isSecret)
}

func testAccResourceParameterCreateWithRules(projName, resourceName, paramName, desc string, isSecret bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "string"
		rule {
			type       = "min_len"
			constraint = "1"
		}
		rule {
			type       = "max_len"
			constraint = "10"
		}
		rule {
			type       = "regex"
			constraint = ".*"
		}
	}
	`, resourceName, projName, paramName, desc, isSecret)
}
