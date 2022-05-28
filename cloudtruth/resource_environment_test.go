package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const envdesc = "Just a description of an environment"
const updateEnvDesc = "A new description of an environment"
const parentEnv = "default"

func TestAccResourceEnvironmentBasic(t *testing.T) {
	createEnvName := fmt.Sprintf("TestEnv-%s", resource.UniqueId())
	updateEnvName := fmt.Sprintf("updated-%s", createEnvName)
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentCreateBasic(createEnvName, envdesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "name", createEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "description", envdesc),
				),
			},
			{
				Config: testAccResourceEnvironmentUpdateBasic(updateEnvName, updateEnvDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "name", updateEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.basic", "description", updateEnvDesc),
				),
			},
		},
	})
}

func testAccResourceEnvironmentCreateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, envName, desc)
}

func testAccResourceEnvironmentUpdateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "basic" {
  		name        = "%s"
  		description = "%s"
	}
	`, envName, desc)
}
