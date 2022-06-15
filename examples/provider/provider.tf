/*
The path convention for providers is:
HOSTNAME/NAMESPACE/TYPE/terraform-provider-TYPE_VERSION_TARGET.zip
In our case that's "registry.terraform.io/cloudtruth/cloudtruth"

This does not look redundant when the company that owns the provider
is not the same as the service the provider is build for e.g.
registry.terraform.io/hashicorp/aws
*/

terraform {
  required_version = ">=0.15.0"
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
      version = "0.0.3"
    }
  }
}

# Recommended - set this via
# export TF_VAR_cloudtruth_api_key
variable "cloudtruth_api_key" {}

provider "cloudtruth" {
  api_key = var.cloudtruth_api_key
}
