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
	createPushProject := "AcceptanceTest"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationName, genericDesc,
					false, false, false, true, true, true, true,
					createRegion, createService, createPushPattern, createPushProject),
				ExpectError: regexp.MustCompile("one of `include_parameters`, `include_secrets`, or `include_templates` must be true"),
				SkipFunc:    isSelfHostedOrStaging,
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
	createPushProject, updatePushProjects := "AcceptanceTest", "PushTestProject"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationName, genericDesc,
					true, false, true, true, true, true, true,
					createRegion, createService, createPushPattern, createPushProject),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource", createPushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "parameters", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "secrets", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "templates", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "coerce", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "force", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "local", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "projects.0", createPushProject),
				),
				SkipFunc: isSelfHostedOrStaging,
			}, {
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationName, genericDesc,
					false, true, false, false, false, false, false,
					createRegion, createService, updatePushPattern, updatePushProjects),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource", updatePushPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "parameters", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "secrets", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "templates", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "coerce", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "force", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "local", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "projects.0", updatePushProjects),
				),
				SkipFunc: isSelfHostedOrStaging,
			}, {
				Config: testAccResourceAWSPushActionBadTag(resourceName, pushActionName, accTestAWSIntegrationName, genericDesc,
					false, true, false, false, false, false, false,
					createRegion, createService, updatePushPattern, createPushProject),
				ExpectError: regexp.MustCompile("did not find the tag"),
				SkipFunc:    isSelfHostedOrStaging,
			},
		},
	})
}

func testAccResourceAWSPushActionBasic(resource, name, intName, desc string, params, secrets, templates, coerce, force, local, dryRun bool, region, service, resourcePattern, projects string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_push_action" "%s" {
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
		region               = "%s"
		service              = "%s"
		resource		     = "%s"
		projects             = ["%s"]
		tags                 = ["default:EpochTime"]
	}
	`, resource, name, intName, desc, params, secrets, templates, coerce, force, local, dryRun, region, service, resourcePattern, projects)
}

func testAccResourceAWSPushActionBadTag(resource, name, intName, desc string, params, secrets, templates, coerce, force, local, dryRun bool, region, service, resourcePattern, projects string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_push_action" "%s" {
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
		region               = "%s"
		service              = "%s"
		resource             = "%s"
		projects             = ["%s"]
		tags                 = ["default:xxx"]
	}
	`, resource, name, intName, desc, params, secrets, templates, coerce, force, local, dryRun, region, service, resourcePattern, projects)
}
