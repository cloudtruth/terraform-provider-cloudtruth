// Fixtures needed for Acceptance tests
// This could be wired up into a daily Atlantis type flow to ensure
// that these resources always exist, not sure it's worthwhile at this point

resource "cloudtruth_project" "acceptance_test" {
  name        = "AcceptanceTest"
  description = "A project used by CI/CD jobs for the CloudTruth Terraform provider"
}

resource "cloudtruth_parameter" "default_regular" {
  name        = "DefaultRegularParam"
  type        = "string"
  description = "Just a description of a parameter"
  project     = cloudtruth_project.acceptance_test.name
}

resource "cloudtruth_parameter_value" "default_regular" {
  parameter_name = cloudtruth_parameter.default_regular.name
  value          = "notreallyasecret"
}

resource "cloudtruth_parameter" "default_secret" {
  name        = "DefaultSecretParam"
  type        = "string"
  secret      = true
  description = "A secret parameter used for acceptance tests"
  project     = cloudtruth_project.acceptance_test.name
}

resource "cloudtruth_parameter_value" "default_secret" {
  parameter_name = cloudtruth_parameter.default_secret.name
  value          = "ultratopsecret"
}

/*resource "cloudtruth_parameter" "default_regular_external" {
  name        = "DefaultRegularExternalParam"
  type        = "string"
  description = "For simplicity/stability, we reference the provider repo's manifest file which is required by the Terraform Registry"
  project     = cloudtruth_project.acceptance_test.name
}

resource "cloudtruth_parameter_value" "default_regular_external" {
  parameter_name = cloudtruth_parameter.default_regular_external.name
  external       = true
  location       = "github://cloudtruth/terraform-provider-cloudtruth/main/terraform-registry-manifest.json"
  filter         = "metadata.protocol_versions[0]"
}*/

resource "cloudtruth_tag" "epoch_time" {
  name        = "EpochTime"
  environment = "default"
  timestamp   = "1970-01-01T05:00:00Z"
}

resource "cloudtruth_template" "basic" {
  name    = "BasicTemplate"
  project = cloudtruth_project.acceptance_test.name
  value   = "Regular Parameter: {{DefaultRegularParam}}, Secret Parameter {{DefaultSecretParam}}"
}

resource "cloudtruth_template" "default" {
  name    = "DefaultTemplate"
  project = cloudtruth_project.acceptance_test.name
  value   = "{{DefaultRegularParam}}"
}


