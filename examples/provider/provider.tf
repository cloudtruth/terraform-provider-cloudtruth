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
