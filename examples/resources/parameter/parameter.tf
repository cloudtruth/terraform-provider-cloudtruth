resource "cloudtruth_parameter" "example" {
  name        = "MyNewParameter"
  description = "This is a sample CloudTruth parameter"
  # environment = "production" # optional, defaults to 'default'
  project = "MyNewProject"
  value   = "some_value"
  secret  = false # optional, defaults to false
  dynamic = true
}
