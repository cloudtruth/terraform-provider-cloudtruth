package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"strconv"
	"testing"
)

func TestAccResourceParameterWithBadRules(t *testing.T) {
	resourceName := "bad_rules"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterDisallowedBooleanRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("the base type boolean does not support the min_len rule type"),
			},
			{
				Config: testAccResourceParameterInvalidIntegerRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("the base type integer does not support the min_len rule type"),
			},
			{
				Config: testAccResourceParameterInvalidStringRules(accTestProject, resourceName,
					fmt.Sprintf("Test-%s", uuid.New().String())),
				ExpectError: regexp.MustCompile("the base type string does not support the min rule type"),
			},
		},
	})
}

func TestAccResourceParameterWithRules(t *testing.T) {
	createStringParamName := fmt.Sprintf("Test-Str-%s", uuid.New().String())
	createIntegerParamName := fmt.Sprintf("Test-Int-%s", uuid.New().String())
	stringResourceName, intResourceName := "string_with_rules", "int_with_rules"
	min, max := 1, 10
	createRegEx := ".*"
	updateMin, updateMax := 0, 11
	updateRegEx := `123.*`
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateStringWithRules(accTestProject, stringResourceName, createStringParamName, paramDesc,
					false, min, max, createRegEx),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "min_len", fmt.Sprint(min)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "max_len", fmt.Sprint(max)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", createRegEx),
				),
			},
			{
				Config: testAccResourceParameterCreateStringWithRules(accTestProject, stringResourceName, createStringParamName, paramDesc,
					false, updateMin, updateMax, updateRegEx),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "min_len", fmt.Sprint(updateMin)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "max_len", fmt.Sprint(updateMax)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", updateRegEx),
				),
			},
			{
				Config: testAccResourceParameterCreateIntegerWithRules(accTestProject, intResourceName, createIntegerParamName, paramDesc,
					false, min, max),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "name", createIntegerParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min", fmt.Sprint(min)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max", fmt.Sprint(max)),
				),
			},
			{
				Config: testAccResourceParameterCreateIntegerWithRules(accTestProject, intResourceName, createIntegerParamName, paramDesc,
					false, min, max),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "name", createIntegerParamName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "description", paramDesc),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "secret",
						strconv.FormatBool(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min", fmt.Sprint(updateMin)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max", fmt.Sprint(updateMax)),
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

func testAccResourceParameterCreateIntegerWithRules(projName, resourceName, paramName, desc string, isSecret bool,
	minLen, maxLen int) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "integer"
		min         = %d
        max         = %d
	}
	`, resourceName, projName, paramName, desc, isSecret, minLen, maxLen)
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
        min_len     = 1
	}
	`, resourceName, projName, paramName)
}

// Integer parameter types can have max|min rules, string types can have max_len|min_len|regex rules
func testAccResourceParameterInvalidStringRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
        min         = 1
	}
	`, resourceName, projName, paramName)
}
