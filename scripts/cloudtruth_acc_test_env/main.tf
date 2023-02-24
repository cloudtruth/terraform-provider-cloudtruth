// Fixtures needed for Acceptance tests

resource "cloudtruth_parameter" "default_regular" {
  name = "DefaultRegularParam"
  type = "string"
  description = "Just a description of a parameter"
}

resource "cloudtruth_parameter_value" "default_regular" {
  parameter_name = cloudtruth_parameter.default_regular.name
  value = "notreallyasecret"
}

resource "cloudtruth_parameter" "default_secret" {
  name = "DefaultSecretParam"
  type = "string"
  secret = true
  description = "A secret parameter used for acceptance tests"
}

resource "cloudtruth_parameter_value" "default_secret" {
  parameter_name = cloudtruth_parameter.default_secret.name
  value = "ultratopsecret"
}

# Need to add the integration externally/by hand for now
resource "cloudtruth_parameter" "default_regular_external" {
  name = "DefaultRegularExternalParam"
  type = "string"
  description = "For simplicity/stability, we reference the provider repo's manifest file which is required by the Terraform Registry"
}

resource "cloudtruth_parameter_value" "default_regular_external" {
  parameter_name = cloudtruth_parameter.default_regular_external.name
  external = true
  location = "github://cloudtruth/terraform-provider-cloudtruth/main/terraform-registry-manifest.json"
  filter = "metadata.protocol_versions[0]"
}
