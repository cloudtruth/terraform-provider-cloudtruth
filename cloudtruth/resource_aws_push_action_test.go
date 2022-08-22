package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceAWSPushActionBasic(t *testing.T) {
	resourceName := "basic"
	pushActionName := fmt.Sprintf("TestAWSPush-%s", uuid.New().String())
	createService := "secretsmanager"
	createRegion := "us-east-1"
	//createPushPattern, updatePushPattern := "/{{ environment }}/{{ project }}/{{ parameter }}", "/some_prefix/{{ environment }}/{{ project }}/{{ parameter }}"
	createPushPattern := "/{{ environment }}/{{ project }}/{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationID, genericDesc, createRegion, createService, createPushPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "create_environments", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "create_projects", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource_pattern", createPushPattern),
				),
			}, /*
				{
					Config: testAccResourceAWSPushActionBasic(resourceName, pushActionName, accTestAWSIntegrationID, genericDesc, false, false,
						createRegion, createService, updatePushPattern, projects),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "name", pushActionName),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "create_environments", fmt.Sprint(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "create_projects", fmt.Sprint(false)),
						resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_push_action.%s", resourceName), "resource_pattern", updatePushPattern),
					),
				},*/
		},
	})
}

func testAccResourceAWSPushActionBasic(resource, name, intID, desc string, region, service, resourcePattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_push_action" "%s" {
  		name                 = "%s"
		integration_id       = "%s"
  		description          = "%s"
		region               = "%s"
		service              = "%s"
		resource_pattern     = "%s"
		projects             = ["AcceptanceTest"]
		tags                 = ["default:EpochTime"]
	}
	`, resource, name, intID, desc, region, service, resourcePattern)
}
