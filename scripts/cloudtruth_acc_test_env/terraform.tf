terraform {
  required_providers {
    cloudtruth = {
      source = "cloudtruth/cloudtruth"
      version = ">= 0.5.2"
    }
  }
}

provider "cloudtruth" {
  domain = "api.staging.cloudtruth.io"
  project = "AcceptanceTest"
}

