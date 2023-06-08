package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccResourceAzurePushActionInvalid(t *testing.T) {
	resourceName := "basic"
	pushActionName := fmt.Sprintf("TestAzurePush-%s", uuid.New().String())
	createPushPattern := "{{ environment }}-{{ project }}-{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAzurePushActionBasic(resourceName, pushActionName, accTestAzureIntegrationName, genericDesc,
					false, false, true, true, true, true, true, createPushPattern),
				ExpectError: regexp.MustCompile("one of `include_parameters` or `include_secrets` must be true"),
				SkipFunc:    isSelfHostedOrStaging,
			},
		},
	})
}

func TestAccResourceAzurePushActionBasic(t *testing.T) {
	resourceName := "basic"
	pushActionName := fmt.Sprintf("TestAzurePush-%s", uuid.New().String())
	createPushPattern, updatePushPattern := "{{ environment }}-{{ project }}-{{ parameter }}", "some-prefix-{{ environment }}-{{ project }}-{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAzurePushActionBasic(resourceName, pushActionName, accTestAzureIntegrationName, genericDesc,
					true, false, true, true, true, true, true, createPushPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "resource", createPushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "parameters", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "secrets", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "templates", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "coerce", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "force", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "local", fmt.Sprint(true)),
				),
				SkipFunc: isSelfHostedOrStaging,
			}, {
				Config: testAccResourceAzurePushActionBasic(resourceName, pushActionName, accTestAzureIntegrationName, genericDesc,
					false, true, true, false, false, false, false, updatePushPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "resource", updatePushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "parameters", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "secrets", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "coerce", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "force", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_push_action.%s", resourceName), "local", fmt.Sprint(false)),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
		},
	})
}

func testAccResourceAzurePushActionBasic(resource, name, intName, desc string, params, secrets, templates, coerce, force, local, dryRun bool, resourcePattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_azure_push_action" "%s" {
  		name                 = "%s"
		integration          = "%s"
  		description          = "%s"
		parameters           = "%t"
		secrets              = "%t"
		templates            = "%t"
		coerce               = "%t"
		force                = "%t"
		local                = "%t"
		dry_run              = "%t"
		resource		     = "%s"
		projects             = ["AcceptanceTest"]
		tags                 = ["default:EpochTime"]
	}
	`, resource, name, intName, desc, params, secrets, templates, coerce, force, local, dryRun, resourcePattern)
}
