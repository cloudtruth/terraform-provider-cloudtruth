terraform {
  required_version = ">=0.13.0"

  required_providers {
    cloudtruth = {
      version = "0.0.1"
    }

    sysdig = {
      source = "sysdiglabs/sysdig"
      version = "0.5.37"
    }
  }
}

provider "cloudtruth" {
  api_version = "v1"
  api_key = "should_use_env_var_instead"
}
