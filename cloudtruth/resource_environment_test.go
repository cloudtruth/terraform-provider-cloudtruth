package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
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

func TestAccResourceEnvForceDelete(t *testing.T) {
	forceDeleteEnvName := fmt.Sprintf("TestDeleteEnv-%s", uuid.New().String())
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		// todo: uncomment when https://github.com/hashicorp/terraform-plugin-sdk/pull/976 has been merged
		// and allows us to confirm that the delete indeed fails, for now we
		//ExpectDestroyError: regexp.MustCompile(`.*cannot be deleted from the CloudTruth provider, you must enable 'force_delete' to allow this deletes.*`),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvForceDeleteDisabled(forceDeleteEnvName, envdesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "name", forceDeleteEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "description", envdesc),
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "force_delete",
						strconv.FormatBool(false)),
				),
			},
			{ // todo: move this to a new Test(Step) when we can validate that the above delete fails
				// keeping it here now ensures that temporary resource cleanup will happen
				Config: testAccResourceEnvForceDeleteEnabled(forceDeleteEnvName, envdesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "name", forceDeleteEnvName),
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "description", envdesc),
					resource.TestCheckResourceAttr("cloudtruth_environment.force_delete", "force_delete",
						strconv.FormatBool(true)),
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
        force_delete = true
	}
	`, envName, desc)
}

func testAccResourceEnvUpdateBasic(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "basic" {
  		name         = "%s"
  		description  = "%s"
        force_delete = true
	}
	`, envName, desc)
}

func testAccResourceEnvForceDeleteDisabled(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "force_delete" {
  		name         = "%s"
  		description  =  "%s"
	}
	`, envName, desc)
}

func testAccResourceEnvForceDeleteEnabled(envName, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "force_delete" {
  		name         = "%s"
  		description  =  "%s"
		force_delete = true
	}
	`, envName, desc)
}

func testAccResourceEnvNested(envOne, envTwo, envThree string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_environment" "env_one" {
  		name         = "%s"
        force_delete = true
	}

	resource "cloudtruth_environment" "env_two" {
  		name         = "%s"
        force_delete = true
  		parent       = cloudtruth_environment.env_one.name
	}

	resource "cloudtruth_environment" "env_three" {
  		name         = "%s"
        force_delete = true
  		parent       = cloudtruth_environment.env_two.name
	}
	`, envOne, envTwo, envThree)
}
