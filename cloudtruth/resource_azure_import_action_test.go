package cloudtruth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceAzureImportActionBasic(t *testing.T) {
	resourceName := "basic"
	importActionName := fmt.Sprintf("TestAzureImport-%s", uuid.New().String())
	createImportPattern, updateImportPattern := "{{ environment }}-{{ project }}-{{ parameter }}", "someprefix-{{ environment }}-{{ project }}-{{ parameter }}"
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAzureImportActionBasic(resourceName, importActionName, accTestAzureIntegrationName, genericDesc, true, true,
					createImportPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "name", importActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "create_environments", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "create_projects", fmt.Sprint(true)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "resource", createImportPattern),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "mode", "pattern"),
				),
			},
			{
				Config: testAccResourceAzureImportActionBasic(resourceName, importActionName, accTestAzureIntegrationName, genericDesc, false, false,
					updateImportPattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "name", importActionName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "create_environments", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "create_projects", fmt.Sprint(false)),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_import_action.%s", resourceName), "resource", updateImportPattern),
				),
			},
		},
	})
}

func testAccResourceAzureImportActionBasic(resource, name, intName, desc string, createEnvs, createProjs bool, resourcePattern string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_azure_import_action" "%s" {
  		name                 = "%s"
		integration          = "%s"
  		description          = "%s"
        create_environments  = "%t"
		create_projects      = "%t"
		resource             = "%s"
		mode                 = "pattern"
	}
	`, resource, name, intName, desc, createEnvs, createProjs, resourcePattern)
}
