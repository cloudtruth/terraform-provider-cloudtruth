resource "cloudtruth_environment" "example" {
  name        = "MyNewEnvironment"
  description = "This is a sample CloudTruth environment"
  force_delete   = true
}
