resource "cloudtruth_parameter" "example" {
  name             = "MyNewParameter"
  description      = "This is a sample CloudTruth parameter"
  environment      = "default"
  project          = "MyCloudTruthProject"
  secret           = false
  type             = "string" # one of string, boolean, integer or custom, defaults to string
  value            = "some_value"
  #evaluate         = true
  #wrap             = false
  # force_delete   = true # set this to allow Terraform to delete project
}
