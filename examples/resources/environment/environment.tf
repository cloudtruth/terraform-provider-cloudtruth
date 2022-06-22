terraform {
  required_version = ">=0.15.0"
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
    }
  }
}

resource "cloudtruth_environment" "example" {
  name         = "MyNewEnvironment"
  description  = "This is a sample CloudTruth environment"
  force_delete = true
}
