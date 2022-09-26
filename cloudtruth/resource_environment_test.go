package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const envdesc = "Just a description of an environment"
const updateEnvDesc = "A new description of an environment"

func TestAccResourceEnvBasic(t *testing.T) {
	createEnvName := fmt.Sprintf("TestEnv-%s", uuid.New().String())
	updateEnvName := fmt.Sprintf("updated-%s", createEnvName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvCreateBasic(createEnvName, envdesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "name", createEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "description", envdesc),
				),
			},
			{
				Config: testAccResourceEnvUpdateBasic(updateEnvName, updateEnvDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "name", updateEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "description", updateEnvDesc),
				),
			},
		},
	})
}

func TestAccResourceEnvNested(t *testing.T) {
	envOneName := fmt.Sprintf("TestEnv1-%s", uuid.New().String())
	envTwoName := fmt.Sprintf("TestEnv2-%s", uuid.New().String())
	envThreeName := fmt.Sprintf("TestEnv3-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvNested(envOneName, envTwoName, envThreeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.env_one", "name", envOneName),
					resource.TestCheckResourceAttr("cloudtruth_environment.env_two", "name", envTwoName),
					resource.TestCheckResourceAttr("cloudtruth_environment.env_three", "name", envThreeName),
				),
			},
		},
	})
}

func testAccResourceEnvCreateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "basic" {
  		name         = "%s"
  		description  = "%s"
	}
	`, envName, desc)
}

func testAccResourceEnvUpdateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "basic" {
  		name         = "%s"
  		description  = "%s"
	}
	`, envName, desc)
}

func testAccResourceEnvNested(envOne, envTwo, envThree string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "env_one" {
  		name         = "%s"
	}

	resource "cloudtruth_environment" "env_two" {
  		name         = "%s"
  		parent       = cloudtruth_environment.env_one.name
	}

	resource "cloudtruth_environment" "env_three" {
  		name         = "%s"
  		parent       = cloudtruth_environment.env_two.name
	}
	`, envOne, envTwo, envThree)
}
