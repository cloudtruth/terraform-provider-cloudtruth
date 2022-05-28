resource "cloudtruth_environment" "example" {
  name             = "MyNewEnvironment"
  description      = "This is a sample CloudTruth environment"
  # parent           = "OptionalParentEnvironment"
  # force_delete     = true # set this to allow Terraform to delete project
}
