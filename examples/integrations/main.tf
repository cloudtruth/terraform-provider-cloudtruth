# Sample .tf for a CloudTruth AWS Role and integrating with an AWS service.
# Apply this to one or more of your AWS accounts to use CloudTruth AWS integrations
# See https://docs.cloudtruth.com/integrations/aws
provider "aws" {
}

module "grant_cloudtruth_access" {
  source                 = "github.com/cloudtruth/terraform-cloudtruth-access"
  role_name              = var.role_name
  external_id            = var.external_id
  services_enabled       = var.services_enabled
  services_write_enabled = var.services_write_enabled
}
