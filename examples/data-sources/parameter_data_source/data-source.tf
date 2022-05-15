# todo: get auth working: env, profile, inline in .tf?
# specify env + projec
# Handle wrapping/unwrapping, need to find out about decryption
# support id and name

data "parameter_data_source" "example" {
  env = "production"
  project = "MyFirstProject"
  wrap = true
}
