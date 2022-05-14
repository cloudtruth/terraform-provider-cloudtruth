# todo: get auth working: env, profile, inline in .tf?
# specify env + projec

data "parameter_data_source" "example" {
  env = "production" # support id and name
  project = "MyFirstProject" # support id and name
  wrap = true # Handle wrapping/unwrapping, need to find out about decryption
}