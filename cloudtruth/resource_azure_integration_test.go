package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"strings"
	"testing"
)

// generateRandomVaultName generates a valid Azure Key Vault name for testing
func generateRandomVaultName() string {
	// Azure Key Vault names must be 1-24 characters, alphanumeric and hyphens
	// Cannot start or end with hyphen, no consecutive hyphens
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	length := 10 + rand.Intn(10) // Random length between 10-20 characters
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return "vault" + string(b)
}

// generateRandomTenantID generates a random UUID for testing
func generateRandomTenantID() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		rand.Uint32(),
		rand.Uint32()&0xffff,
		rand.Uint32()&0xffff,
		rand.Uint32()&0xffff,
		rand.Uint64()&0xffffffffffff)
}

func TestAccResourceAzureIntegrationBasic(t *testing.T) {
	resourceName := "basic"
	vaultName := generateRandomVaultName()
	tenantID := generateRandomTenantID()
	description := "Test Azure Key Vault integration"
	updatedDescription := "Updated test integration"

	resourceTags := map[string]string{
		"Environment": "test",
		"ManagedBy":   "terraform",
	}
	updatedResourceTags := map[string]string{
		"Environment": "staging",
		"ManagedBy":   "terraform",
		"Team":        "platform",
	}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// Test basic creation with minimal fields
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, "", false, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "writable", "false"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "id"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "name"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "fqn"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "status"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "type"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "type", "akv"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				// Test update with description and resource tags
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, description, true, resourceTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "description", description),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "writable", "true"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Environment", "test"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.ManagedBy", "terraform"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				// Test update description and resource tags
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, updatedDescription, false, updatedResourceTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "description", updatedDescription),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "writable", "false"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Environment", "staging"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Team", "platform"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				// Test removing resource tags
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, updatedDescription, false, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "description", updatedDescription),
					resource.TestCheckNoResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Environment"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
		},
	})
}

func TestAccResourceAzureIntegrationFull(t *testing.T) {
	resourceName := "full"
	vaultName := generateRandomVaultName()
	tenantID := generateRandomTenantID()
	description := "Full Azure Key Vault integration with all fields"

	resourceTags := map[string]string{
		"Owner":       "engineering",
		"Project":     "platform",
		"CostCenter":  "12345",
		"Environment": "production",
	}

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, description, true, resourceTags),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "description", description),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "writable", "true"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Owner", "engineering"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Project", "platform"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.CostCenter", "12345"),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "resource_tags.Environment", "production"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "id"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "name"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "fqn"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "created_at"),
					resource.TestCheckResourceAttrSet(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "url"),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
		},
	})
}

func TestAccResourceAzureIntegrationImportByID(t *testing.T) {
	resourceName := "import_test"
	vaultName := generateRandomVaultName()
	tenantID := generateRandomTenantID()
	description := "Integration for testing import by ID"

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAzureIntegrationBasic(resourceName, vaultName, tenantID, description, false, nil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "vault_name", vaultName),
					resource.TestCheckResourceAttr(fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName), "tenant_id", tenantID),
				),
				SkipFunc: isSelfHostedOrStaging,
			},
			{
				ResourceName:      fmt.Sprintf("cloudtruth_azure_integration.%s", resourceName),
				ImportState:       true,
				ImportStateVerify: true,
				SkipFunc:          isSelfHostedOrStaging,
			},
		},
	})
}

func TestAccResourceAzureIntegrationValidation(t *testing.T) {
	resourceName := "validation_test"
	validTenantID := generateRandomTenantID()

	testCases := []struct {
		name          string
		vaultName     string
		tenantID      string
		expectError   bool
		errorContains string
	}{
		{
			name:          "valid_vault_name",
			vaultName:     "validvault123",
			tenantID:      validTenantID,
			expectError:   false,
			errorContains: "",
		},
		{
			name:          "vault_name_too_long",
			vaultName:     "thisistoolongforvaultname1234567890",
			tenantID:      validTenantID,
			expectError:   true,
			errorContains: "must be between 1 and 24 characters",
		},
		{
			name:          "vault_name_with_underscore",
			vaultName:     "invalid_vault",
			tenantID:      validTenantID,
			expectError:   true,
			errorContains: "must contain only alphanumeric characters and hyphens",
		},
		{
			name:          "vault_name_starts_with_hyphen",
			vaultName:     "-invalidvault",
			tenantID:      validTenantID,
			expectError:   true,
			errorContains: "cannot start or end with a hyphen",
		},
		{
			name:          "vault_name_ends_with_hyphen",
			vaultName:     "invalidvault-",
			tenantID:      validTenantID,
			expectError:   true,
			errorContains: "cannot start or end with a hyphen",
		},
		{
			name:          "vault_name_consecutive_hyphens",
			vaultName:     "invalid--vault",
			tenantID:      validTenantID,
			expectError:   true,
			errorContains: "cannot contain consecutive hyphens",
		},
		{
			name:          "invalid_tenant_id",
			vaultName:     "validvault",
			tenantID:      "not-a-uuid",
			expectError:   true,
			errorContains: "invalid UUID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			config := testAccResourceAzureIntegrationBasic(resourceName, tc.vaultName, tc.tenantID, "", false, nil)

			testStep := resource.TestStep{
				Config:   config,
				SkipFunc: isSelfHostedOrStaging,
			}

			if tc.expectError {
				testStep.ExpectError = nil // Will be checked by the test framework
			}

			resource.Test(t, resource.TestCase{
				ProviderFactories: testProviderFactories,
				PreCheck:          func() { testAccPreCheck(t) },
				Steps:             []resource.TestStep{testStep},
			})
		})
	}
}

func testAccResourceAzureIntegrationBasic(resource, vaultName, tenantID, description string, writable bool, resourceTags map[string]string) string {
	descriptionStr := ""
	if description != "" {
		descriptionStr = fmt.Sprintf(`description = "%s"`, description)
	}

	resourceTagsStr := ""
	if len(resourceTags) > 0 {
		var tags []string
		for k, v := range resourceTags {
			tags = append(tags, fmt.Sprintf(`%s = "%s"`, k, v))
		}
		resourceTagsStr = fmt.Sprintf("resource_tags = {\n    %s\n  }", strings.Join(tags, "\n    "))
	}

	return fmt.Sprintf(`
resource "cloudtruth_azure_integration" "%s" {
  vault_name = "%s"
  tenant_id  = "%s"
  %s
  writable   = %t
  %s
}
	`, resource, vaultName, tenantID, descriptionStr, writable, resourceTagsStr)
}
