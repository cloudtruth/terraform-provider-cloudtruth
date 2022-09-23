terraform init
# Explicit syntax which specifies the environment
terraform import cloudtruth_parameter_value.example PROJECT_NAME.ENVIRONMENT_NAME.PARAMETER_ID.PARAMETER_VALUE_ID

# When no environment is specified, the import defaults to the 'default' environment
terraform import cloudtruth_parameter_value.example PROJECT_NAME.PARAMETER_ID.PARAMETER_VALUE_ID

