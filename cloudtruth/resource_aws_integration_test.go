package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
)

func TestAccResourceAWSIntegrationBasic(t *testing.T) {
	resourceName := "basic"
	accountID, updateAccountID := fmt.Sprintf("%012d",
		rand.Int63n(1e12)), fmt.Sprintf("%012d", rand.Int63n(1e12))
	role, updateRole := "SomeAWSRole", "SomeOtherRole"
	kmsKeyID, updateKMSKeyID := "10bfd420-0369-4be3-8546-20f21d7bf2a7", "XXXfd420-0369-4be3-8546-20f21d7bf2a7"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSIntegrationBasic(resourceName, accountID, role, kmsKeyID, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "account_id", accountID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "role", role),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "kms_key_id", kmsKeyID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "writable", fmt.Sprint(true)),
				),
			},
			{
				Config: testAccResourceAWSIntegrationBasic(resourceName, updateAccountID, updateRole, updateKMSKeyID, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "account_id", updateAccountID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "role", updateRole),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "kms_key_id", updateKMSKeyID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_integration.%s", resourceName), "writable", fmt.Sprint(false)),
				),
			},
		},
	})
}

func testAccResourceAWSIntegrationBasic(resource, accountID, role, kmsKeyID string, writable bool) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_integration" "%s" {
  		account_id           = "%s"
		role                 = "%s"
  		kms_key_id           = "%s"
        writable             = "%t"
		aws_enabled_regions  = ["us-east-1", "us-east-2"]
		aws_enabled_services = ["ssm", "secretsmanager"]
	}
	`, resource, accountID, role, kmsKeyID, writable)
}
