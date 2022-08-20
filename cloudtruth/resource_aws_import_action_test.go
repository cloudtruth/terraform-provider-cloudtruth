package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceAWSImportActionBasic(t *testing.T) {
	resourceName := "basic"
	importActionName := fmt.Sprintf("TestAWSImport-%s", uuid.New().String())

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSImportActionBasic(resourceName, importActionName, accTestAWSIntegrationID, genericDesc, true, true,
					defaultAWSRegion, "secretsmanager", defaultImportPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "name", importActionName),
				),
			},
		},
	})
}

func testAccResourceAWSImportActionBasic(resource, name, intID, desc string, createEnvs, createProjs bool, region, service, resource_pattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_import_action" "%s" {
  		name                 = "%s"
		integration_id       = "%s"
  		description          = "%s"
        create_environments  = "%t"
		create_projects      = "%t"
		region               = "%s"
		service              = "%s"
		resource_pattern     = "%s"
	}
	`, resource, name, intID, desc, createEnvs, createProjs, region, service, resource_pattern)
}
