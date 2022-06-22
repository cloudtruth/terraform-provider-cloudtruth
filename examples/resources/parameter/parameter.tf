terraform {
  required_version = ">=0.15.0"
  required_providers {
    cloudtruth = {
      source = "cloudtruth/cloudtruth"
    }
  }
}

resource "cloudtruth_parameter" "example" {
  name        = "MyNewParameter"
  description = "This is a sample CloudTruth parameter"
  # environment = "production" # optional, defaults to 'default'
  project = "AcceptanceTest"
  value   = "some_value"
  secret  = false # optional, defaults to false
  dynamic = true
}
