terraform {
  required_version = ">=0.15.0"
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

data "cloudtruth_parameter" "param_example" {
  project     = "AcceptanceTest"
  environment = "default"
  name        = "DefaultRegularParam"
}

# This data source collects a map of all values for the parameter
# in all environments where it's defined
data "cloudtruth_parameters" "params_example" {
  project     = "AcceptanceTest"
  name        = "DefaultRegularParam"
}

# Example outputs
output "parameter_output" {
  value = data.cloudtruth_parameter.param_example
}

output "parameters_output" {
  value = data.cloudtruth_parameters.params_example
}
