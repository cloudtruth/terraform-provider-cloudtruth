resource "cloudtruth_parameter" "example" {
  name             = "MyNewParameter"
  description      = "This is a sample CloudTruth parameter"
  environment      = "default"
  project          = "MyCloudTruthProject"
  secret           = false
  type             = "string" # one of string, boolean, integer or custom, defaults to string
  # defaults to empty string, when a string and 0/false when integer/boolean
  # if you need to inject this value from elsewhere, omit value and add an ignore_changes lifecycle
  # rule so that Terraform does not attempt to overwrite this value
  value            = "some_value"
  #evaluate         = true
  #wrap             = false
  # force_delete   = true # set this to allow Terraform to delete project
}
