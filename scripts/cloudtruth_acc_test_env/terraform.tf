terraform {
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
      version = ">= 0.5.2"
    }
  }
}

provider "cloudtruth" {
  domain   = "api.cloudtruth.local:8080"
  project  = "AcceptanceTest"
  protocol = "http"
}

