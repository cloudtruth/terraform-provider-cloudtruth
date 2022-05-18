terraform {
  required_version = ">=0.13.0"

  required_providers {
    cloudtruth = {
      source = "cloudtruth/cloudtruth"
      version = "0.0.1"
    }
  }
}

# Recommended - set this via
# export TF_VAR_cloudtruth_api_key
variable cloudtruth_api_key{}

provider "cloudtruth" {
  api_key = var.cloudtruth_api_key
}

data "cloudtruth_parameter" "example" {
  name = "first secret"
  env = "production"
  project = "c3e6f8e0-3323-44fd-8760-39998e5f2610"
}

# Note: this is just an example, you would not want to output
# a CloudTruth parameter if it is used to store a secret
output "parameter_output" {
  value = data.cloudtruth_parameter.example
}
