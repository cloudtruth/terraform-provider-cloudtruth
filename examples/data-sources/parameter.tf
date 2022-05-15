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
  // todo: version support
  // api_version = "v1"
  api_key = "should_use_CLOUDTRUTH_API_KEY_env_var_instead"
}

data "cloudtruth_parameter" "example" {
  name = "first secret"
  env = "production"

  # todo: support project name in addition to id
  # also add a project data source eventually
  project = "c3e6f8e0-3323-44fd-8760-39998e5f2610"
  # project = "MyFirstProject"

  # todo: determine what other search/filter/evaluation/masking
  # parameters to support
}
