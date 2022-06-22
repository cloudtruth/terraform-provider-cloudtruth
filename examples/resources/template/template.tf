terraform {
  required_version = ">=0.15.0"
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
    }
  }
}

resource "cloudtruth_template" "example" {
  name        = "MyNewTemplate"
  description = "This is a sample CloudTruth Template"
  project     = "MyNewProject"
  value       = "some_var={{DefaultRegularParam}}"
}
