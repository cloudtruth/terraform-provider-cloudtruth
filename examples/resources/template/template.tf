resource "cloudtruth_parameter" "example" {
  name             = "MyNewTemplate"
  description      = "This is a sample CloudTruth Template"
  environment      = "default"
  project          = "MyCloudTruthProject"
  # todo: other properties
  # force_delete   = true # set this to allow Terraform to delete project
}
