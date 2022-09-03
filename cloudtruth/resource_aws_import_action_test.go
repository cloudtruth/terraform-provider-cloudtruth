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
	createService := "secretsmanager"
	createRegion := "us-east-1"
	createImportPattern, updateImportPattern := "/{{ environment }}/{{ project }}/{{ parameter }}", "/some_prefix/{{ environment }}/{{ project }}/{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSImportActionBasic(resourceName, importActionName, accTestAWSIntegrationName, genericDesc, true, true,
					createRegion, createService, createImportPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "name", importActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "create_environments", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "create_projects", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "region", createRegion),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "service", createService),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "resource", createImportPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "mode", "pattern"),
				),
			},
			{
				Config: testAccResourceAWSImportActionBasic(resourceName, importActionName, accTestAWSIntegrationName, genericDesc, true, true,
					createRegion, createService, updateImportPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "name", importActionName),
					// todo: uncomment and set corresponding input values to false when the backend bug behind this is fixed
					//resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "create_environments", fmt.Sprint(false)),
					//resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "create_projects", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_aws_import_action.%s", resourceName), "resource", updateImportPattern),
				),
			},
		},
	})
}

func testAccResourceAWSImportActionBasic(resource, name, intName, desc string, createEnvs, createProjs bool, region, service, resourcePattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_aws_import_action" "%s" {
  		name                 = "%s"
		integration          = "%s"
  		description          = "%s"
        create_environments  = "%t"
		create_projects      = "%t"
		region               = "%s"
		service              = "%s"
		resource             = "%s"
		mode                 = "pattern"
	}
	`, resource, name, intName, desc, createEnvs, createProjs, region, service, resourcePattern)
}
