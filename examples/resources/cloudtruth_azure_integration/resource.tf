# Basic Azure Key Vault integration (read-only)
resource "cloudtruth_azure_integration" "example" {
  vault_name  = "mykeyvault"
  tenant_id   = "12345678-1234-1234-1234-123456789abc"
  description = "Production Azure Key Vault integration"
}

# Azure Key Vault integration with write access and resource tags
resource "cloudtruth_azure_integration" "writable" {
  vault_name  = "mykeyvault"
  tenant_id   = "12345678-1234-1234-1234-123456789abc"
  description = "Writable integration for syncing secrets to Azure"
  writable    = true

  resource_tags = {
    Environment = "production"
    ManagedBy   = "terraform"
    Team        = "platform"
  }
}

# Use the integration with import and push actions
resource "cloudtruth_azure_import_action" "import" {
  integration = cloudtruth_azure_integration.example.name  # Format: "vault_name@tenant_id"
  name        = "Import Azure secrets"
  description = "Import secrets from Azure Key Vault"

  resource             = "{{environment}}-{{project}}-{{parameter}}"
  create_environments  = true
  create_projects      = true
}

resource "cloudtruth_azure_push_action" "push" {
  integration = cloudtruth_azure_integration.writable.name
  name        = "Push to Azure"
  description = "Push CloudTruth parameters to Azure Key Vault"

  resource           = "{{project}}-{{environment}}-{{parameter}}"
  coerce_parameters  = true
  include_parameters = true
  include_secrets    = true
  include_templates  = false

  projects = [cloudtruth_project.example.id]
}
