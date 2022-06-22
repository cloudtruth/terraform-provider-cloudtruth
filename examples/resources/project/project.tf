terraform {
  required_version = ">=0.15.0"
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
    }
  }
}

resource "cloudtruth_project" "example" {
  name         = "MyNewProject"
  description  = "some project"
  force_delete = true # set this to allow Terraform to delete project
}
