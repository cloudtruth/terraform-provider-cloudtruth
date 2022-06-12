resource "cloudtruth_project" "example" {
  name         = "MyNewProject"
  description  = "Thisfsaf"
  force_delete = true # set this to allow Terraform to delete project
}
