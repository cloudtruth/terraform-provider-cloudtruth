package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

/*
	func TestAccResourceParameterWithRules(t *testing.T) {
		createStringParamName := fmt.Sprintf("Test-Str-%s", uuid.New().String())
		createIntegerParamName := fmt.Sprintf("Test-Int-%s", uuid.New().String())
		stringResourceName, intResourceName := "string_with_rules", "int_with_rules"
		minVal, maxVal := 1, 10
		createRegEx := ".*"
		updateMin, updateMax := 0, 11
		updateRegEx := `123.*`
		resource.Test(t, resource.TestCase{
			ProviderFactories: testProviderFactories,
			PreCheck:          func() { testAccPreCheck(t) },
			Steps: []resource.TestStep{
				{
					Config: testAccResourceParameterCreateStringWithRules(accTestProject, stringResourceName, createStringParamName, paramDesc,
						false, minVal, maxVal, createRegEx),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "min", fmt.Sprint(minVal)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "max", fmt.Sprint(maxVal)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", createRegEx),
					),
				},
				{
					Config: testAccResourceParameterCreateStringWithRules(accTestProject, stringResourceName, createStringParamName, updateParamDesc,
						false, updateMin, updateMax, updateRegEx),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", updateParamDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret", fmt.Sprint(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "min", fmt.Sprint(updateMin)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "max", fmt.Sprint(updateMax)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", updateRegEx),
					),
				},
				{
					Config: testAccResourceParameterCreateIntegerWithRules(accTestProject, intResourceName, createIntegerParamName, paramDesc,
						false, minVal, maxVal),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "name", createIntegerParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min", fmt.Sprint(minVal)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max", fmt.Sprint(maxVal)),
					),
				},
				{
					Config: testAccResourceParameterCreateIntegerWithRules(accTestProject, intResourceName, createIntegerParamName, paramDesc,
						false, updateMin, updateMax),
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

	func TestAccResourceStringParameterAddRemoveRule(t *testing.T) {
		createStringParamName := fmt.Sprintf("Test-Str-%s", uuid.New().String())
		stringResourceName := "string_with_rule"
		createRegEx := ".*"
		resource.Test(t, resource.TestCase{
			ProviderFactories: testProviderFactories,
			PreCheck:          func() { testAccPreCheck(t) },
			Steps: []resource.TestStep{
				{
					Config: testAccResourceParameterCreateStringWithOneRule(accTestProject, stringResourceName, createStringParamName, paramDesc,
						false, createRegEx),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", createRegEx),
					),
				},
				{
					Config: testAccResourceParameterCreateStringWithNoRules(accTestProject, stringResourceName, createStringParamName, paramDesc,
						false),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "name", createStringParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex", ""),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", stringResourceName), "regex_id", ""),
					),
				},
			},
		})
	}

	func TestAccResourceIntParameterAddRemoveRule(t *testing.T) {
		createIntParamName := fmt.Sprintf("Test-Int-%s", uuid.New().String())
		intResourceName := "int_with_rule"
		minVal, maxVal := 0, 1000
		resource.Test(t, resource.TestCase{
			ProviderFactories: testProviderFactories,
			PreCheck:          func() { testAccPreCheck(t) },
			Steps: []resource.TestStep{
				{
					Config: testAccResourceParameterCreateIntWithRules(accTestProject, intResourceName, createIntParamName, paramDesc,
						false, minVal, maxVal),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "name", createIntParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min", fmt.Sprint(minVal)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max", fmt.Sprint(maxVal)),
					),
				},
				{ // Remove the max and min rules, their corresponding properties should be set to ""
					Config: testAccResourceParameterCreateIntWithNoRules(accTestProject, intResourceName, createIntParamName, paramDesc,
						false),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "name", createIntParamName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "description", paramDesc),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "secret",
							strconv.FormatBool(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min", ""),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max", ""),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "min_id", ""),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_parameter.%s", intResourceName), "max_id", ""),
					),
				},
			},
		})
	}
*/
func TestAccResourceParameterBasic(t *testing.T) {
	createParamName := fmt.Sprintf("Test-%s", uuid.New().String())
	resourceName := "basic"
	fullResourceName := fmt.Sprintf("cloudtruth_parameter.%s", resourceName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceParameterCreateBasic(accTestProject, resourceName, createParamName, paramDesc, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "description", paramDesc),
					resource.TestCheckResourceAttr(fullResourceName, "secret",
						strconv.FormatBool(true)),
				),
			},
			{ // Update test, changing the description and the secret toggle
				Config: testAccResourceParameterCreateBasic(accTestProject, resourceName, createParamName, updateParamDesc, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fullResourceName, "name", createParamName),
					resource.TestCheckResourceAttr(fullResourceName, "description", updateParamDesc),
					resource.TestCheckResourceAttr(fullResourceName, "secret",
						strconv.FormatBool(true)),
				),
			},
			{
				ResourceName:      fullResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: paramOrTemplateStateIDFunc(fullResourceName, accTestProject),
			},
		},
	})
}

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
					ExpectError: regexp.MustCompile("the base type 'boolean' does not support rules"),
				},
				{
					Config: testAccResourceParameterInvalidIntegerRules(accTestProject, resourceName,
						fmt.Sprintf("Test-%s", uuid.New().String())),
					ExpectError: regexp.MustCompile("the base type 'integer' does not support the regex rule type"),
				},
			},
		})
	}
*/
func testAccResourceParameterCreateBasic(projName, resourceName, paramName, desc string, isSecret bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "string"
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
		min         = "%d"
        max         = "%d"
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
		min         = "%d"
        max         = "%d"
	    regex       = "%s"
	}
	`, resourceName, projName, paramName, desc, isSecret, minLen, maxLen, regEx)
}

func testAccResourceParameterCreateStringWithOneRule(projName, resourceName, paramName, desc string, isSecret bool,
	regEx string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "string"
	    regex       = "%s"
	}
	`, resourceName, projName, paramName, desc, isSecret, regEx)
}

func testAccResourceParameterCreateStringWithNoRules(projName, resourceName, paramName, desc string, isSecret bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "string"
	}
	`, resourceName, projName, paramName, desc, isSecret)
}

func testAccResourceParameterCreateIntWithRules(projName, resourceName, paramName, desc string, isSecret bool,
	min, max int) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "integer"
		min         = "%d"
	    max         = "%d"

	}
	`, resourceName, projName, paramName, desc, isSecret, min, max)
}

func testAccResourceParameterCreateIntWithNoRules(projName, resourceName, paramName, desc string, isSecret bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
  		description = "%s"
		secret      = "%t"
		type        = "integer"
	}
	`, resourceName, projName, paramName, desc, isSecret)
}

// Rules are not valid with boolean types
func testAccResourceParameterDisallowedBooleanRules(projName, resourceName, paramName string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_parameter" "%s" {
		project     = "%s"
  		name        = "%s"
		type        = "boolean"
        min         = "123"
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
        regex       = ".*"
	}
	`, resourceName, projName, paramName)
}
