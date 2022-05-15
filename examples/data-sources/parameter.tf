terraform {
  required_version = ">=0.13.0"

  required_providers {
    cloudtruth = {
      source = "cloudtruth/cloudtruth"
      version = "0.0.1"
    }
  }
}

provider "cloudtruth" {
  //api_version = "v1"
  api_key = "should_use_CLOUDTRUTH_API_KEY_env_var_instead"
}

data "cloudtruth_parameter" "example" {
  env = "production"
  project = "MyFirstProject"
  #wrap = true
}
