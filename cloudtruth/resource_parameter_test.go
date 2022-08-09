package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

/*
func TestAccResourceParameterWithBadRules(t *testing.T) {
	resourceName := "bad_rules"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterDisallowedBooleanRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("it does not support rules"),
			},
			{
				Config: testAccResourceParameterTooManyIntegerRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("it accepts no more than two rules"),
			},
			{
				Config: testAccResourceParameterTooManyStringRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("it accepts no more than three rules"),
			},
		},
	})
}
*/
func TestAccResourceParameterWithRules(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "with_rules"
	createMinLen, createMaxLen := 1, 10
	createRegEx := ".*"
	updateMinLen, updateMaxLen := 0, 11
	updateRegEx := `\\d+`
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateStringWithRules(accTestProject, resourceName, createParamName, paramDesc,
					false, createMinLen, createMaxLen, createRegEx),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "min_len", fmt.Sprint(createMinLen)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "max_len", fmt.Sprint(createMaxLen)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "regex", createRegEx),
				),
			},
			{
				Config: testAccResourceParameterCreateStringWithRules(accTestProject, resourceName, createParamName, paramDesc,
					false, updateMinLen, updateMaxLen, updateRegEx),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "name", createParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "min_len", fmt.Sprint(updateMinLen)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "max_len", fmt.Sprint(updateMaxLen)),
					//resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", resourceName), "rule.2.constraint", updateRegEx),
				),
			},
		},
	})
}

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

func testAccResourceParameterCreateStringWithRules(projName, resourceName, paramName, desc string, isSecret bool,
	minLen, maxLen int, regEx string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "string"
		min_len     = %d
        max_len     = %d
	    regex       = "%s"
	}
	`, resourceName, projName, paramName, desc, isSecret, minLen, maxLen, regEx)
}

/*
// Can't have even one rule with boolean types
func testAccResourceParameterDisallowedBooleanRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
		type        = "boolean"
        min_len     = 123
	}
	`, resourceName, projName, paramName)
}

// Invalid types for integer rules, they only accept max and min
func testAccResourceParameterInvalidIntegerRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
		type        = "integer"
	}
	rule {
		type       = "min_len"
		constraint = "1"
	}
	rule {
		type       = "max_len"
		constraint = "12"
	}
	`, resourceName, projName, paramName)
}*/

/*
// There can be no more than two rules per parameter
func testAccResourceParameterTooManyIntegerRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
		type        = "integer"
		rule {
			type       = "min"
			constraint = "1"
		}
		rule {
			type       = "max"
			constraint = "12"
		}
		rule {
			type       = "max"
			constraint = "123"
		}
	}
	`, resourceName, projName, paramName)
}

// Integer parameter types can have max|min rules, string types can have max_len|min_len|regex rules
func testAccResourceParameterInvalidStringRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
	}
	rule {
		type       = "min"
		constraint = "1"
	}
	rule {
		type       = "max"
		constraint = "12"
	}
	`, resourceName, projName, paramName)
}

// There can be no more than two rules per parameter
func testAccResourceParameterTooManyStringRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
		rule {
			type       = "min_len"
			constraint = "1"
		}
		rule {
			type       = "max_len"
			constraint = "12"
		}
		rule {
			type       = "regex"
			constraint = ".*"
		}
		rule {
			type       = ""
			constraint = "\\s+"
		}
	}
	`, resourceName, projName, paramName)
}
*/
