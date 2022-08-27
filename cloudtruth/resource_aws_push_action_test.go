package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"testing"
)

func TestAccResourceAWSPushActionInvalid(t *testing.T) {
	resourceName := "basic"
	pushActionName := fmt.Sprintf("TestAWSPush-%s", uuid.New().String())
	createService := "secretsmanager"
	createRegion := "us-east-1"
	createPushPattern := "/{{ environment }}/{{ project }}/{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationID, genericDesc,
					false, false, true, true, true, true, true,
					createRegion, createService, createPushPattern),
				ExpectError: regexp.MustCompile("one of `include_parameters` or `include_secrets` must be true"),
			},
		},
	})
}

func TestAccResourceAWSPushActionBasic(t *testing.T) {
	resourceName := "basic"
	pushActionName := fmt.Sprintf("TestAWSPush-%s", uuid.New().String())
	createService := "secretsmanager"
	createRegion := "us-east-1"
	createPushPattern, updatePushPattern := "/{{ environment }}/{{ project }}/{{ parameter }}", "/some_prefix/{{ environment }}/{{ project }}/{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationID, genericDesc,
					true, false, true, true, true, true, true,
					createRegion, createService, createPushPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource_pattern", createPushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "parameters", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "secrets", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "templates", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "coerce", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "force", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "local", fmt.Sprint(true)),
				),
			}, {
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationID, genericDesc,
					false, true, false, false, false, false, false,
					createRegion, createService, updatePushPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource_pattern", updatePushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "parameters", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "secrets", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "templates", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "coerce", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "force", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "local", fmt.Sprint(false)),
				),
			},
		},
	})
}

func testAccResourceAWSPushActionBasic(resource, name, intID, desc string, params, secrets, templates, coerce, force, local, dryRun bool, region, service, resourcePattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_push_action" "%s" {
  		name                 = "%s"
		integration_id       = "%s"
  		description          = "%s"
		parameters           = "%t"
		secrets              = "%t"
		templates            = "%t"
		coerce               = "%t"
		force                = "%t"
		local                = "%t"
		dry_run              = "%t"
		region               = "%s"
		service              = "%s"
		resource_pattern     = "%s"
		projects             = ["AcceptanceTest"]
		tags                 = ["default:EpochTime"]
	}
	`, resource, name, intID, desc, params, secrets, templates, coerce, force, local, dryRun, region, service, resourcePattern)
}
