package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"strings"
	"testing"
)

func TestAccResourceAWSIntegrationBasic(t *testing.T) {
	resourceName := "basic"
	accountID, updateAccountID := fmt.Sprintf("%012d",
		rand.Int63n(1e12)), fmt.Sprintf("%012d", rand.Int63n(1e12))
	role, updateRole := "SomeAWSRole", "SomeOtherRole"
	kmsKeyID, updateKMSKeyID := "10bfd420-0369-4be3-8546-20f21d7bf2a7", "XXXfd420-0369-4be3-8546-20f21d7bf2a7"
	resourceTags := map[string]string{
		"tag1": "tag1Val",
	}
	updateResourceTags := map[string]string{
		"tag1": "tag1ValUpdated",
		"tag2": "tag1Val",
	}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSIntegrationBasic(resourceName, accountID, role, kmsKeyID, true, resourceTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "account_id", accountID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "role", role),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "kms_key_id", kmsKeyID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "writable", fmt.Sprint(true)),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "resource_tags.tag1"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "external_id"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				Config: testAccResourceAWSIntegrationBasic(resourceName, updateAccountID, updateRole, updateKMSKeyID, false, updateResourceTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "account_id", updateAccountID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "role", updateRole),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "kms_key_id", updateKMSKeyID),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "resource_tags.tag1"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "resource_tags.tag2"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "writable", fmt.Sprint(false)),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				Config: testAccResourceAWSIntegrationBasic(resourceName, updateAccountID, updateRole, updateKMSKeyID, false, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "account_id", updateAccountID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "role", updateRole),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "kms_key_id", updateKMSKeyID),
					resource.TestCheckNoResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "resource_tags.tag1"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "writable", fmt.Sprint(false)),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
		},
	})
}

func testAccResourceAWSIntegrationBasic(resource, accountID, role, kmsKeyID string, writable bool, resourceTags map[string]string) string {
	resourceTagsStr := ""
	if len(resourceTags) > 0 {
		var tags []string
		for k, v := range resourceTags {
			tags = append(tags, fmt.Sprintf(`%s = "%s"`, k, v))
		}
		resourceTagsStr = fmt.Sprintf("resource_tags = {\n    %s\n  }\n", strings.Join(tags, "\n    "))
	} else {
		resourceTagsStr = ""
	}
	return fmt.Sprintf(`
	resource "cloudtruth_aws_integration" "%s" {
  		account_id           = "%s"
		role                 = "%s"
  		kms_key_id           = "%s"
        writable             = "%t"
		aws_enabled_regions  = ["us-east-2", "us-east-1"]
		aws_enabled_services = ["secretsmanager", "ssm"]
		%s
	}
	`, resource, accountID, role, kmsKeyID, writable, resourceTagsStr)
}
