resource "cloudtruth_project" "example" {
  name             = "MyNewProject"
  description      = "This is a sample cloudtruthproject"
  # parent_project = "OptionalParentProject"
  force_delete     = true # set this to allow Terraform to delete project
}