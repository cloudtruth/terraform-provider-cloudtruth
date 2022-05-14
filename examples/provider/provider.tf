# The path convention for providers is:
# HOSTNAME/NAMESPACE/TYPE/terraform-provider-TYPE_VERSION_TARGET.zip
# In our case: registry.terraform.io/cloudtruth/cloudtruth
# This does not look redundant when the company who owns the provider
# is not the same as the service the provider is build for e.g.
# registry.terraform.io/hashicorp/aws
terraform {
  required_version = ">=0.13.0"

  required_providers {
    cloudtruth = {
      source = "cloudtruth/cloudtruth"
      version = "0.0.1"
    }

    sysdig = {
      source = "sysdiglabs/sysdig"
      version = "0.5.37"
    }
  }
}

provider "cloudtruth" {
  api_version = "v1"
  api_key = "should_use_env_var_instead"
}
