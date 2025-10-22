# CloudTruth Terraform Provider

[![Tests](https://github.com/cloudtruth/terraform-provider-cloudtruth/actions/workflows/test.yml/badge.svg)](https://github.com/cloudtruth/terraform-provider-cloudtruth/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/cloudtruth/terraform-provider-cloudtruth)](https://github.com/cloudtruth/terraform-provider-cloudtruth/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

The official [Terraform](https://www.terraform.io/) provider for [CloudTruth](https://cloudtruth.com/). Manage your CloudTruth configuration as code, enabling infrastructure as code practices for your application configuration and secrets management.

## Features

- **Projects & Environments**: Organize configuration with hierarchical projects and environments
- **Parameters & Values**: Manage configuration parameters with environment-specific values
- **Templates**: Create dynamic configuration templates using CloudTruth parameters
- **Type Validation**: Define custom parameter types with validation rules
- **Access Control**: Manage user groups and access grants
- **Integrations**: Configure AWS and Azure Key Vault integrations for secrets sync
- **Tags**: Version your configuration with environment tags

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.3
- [Go](https://golang.org/doc/install) >= 1.21 (for development)
- [CloudTruth Account](https://app.cloudtruth.com/) and API key

## Installation

### Terraform Registry (Recommended)

The provider is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudtruth/cloudtruth/latest). Add it to your Terraform configuration:

```hcl
terraform {
  required_providers {
    cloudtruth = {
      source  = "cloudtruth/cloudtruth"
      version = "~> 0.9.0"  # Check for latest version
    }
  }
}

provider "cloudtruth" {
  api_key = var.cloudtruth_api_key  # Or use CLOUDTRUTH_API_KEY env var
}
```

### Local Installation (Development)

```bash
git clone https://github.com/cloudtruth/terraform-provider-cloudtruth.git
cd terraform-provider-cloudtruth
make build
```

## Quick Start

### 1. Get Your API Key

Generate an API key from the [CloudTruth dashboard](https://app.cloudtruth.com/).

### 2. Configure the Provider

```hcl
# Set via environment variable (recommended):
# export CLOUDTRUTH_API_KEY=your-api-key-here

provider "cloudtruth" {
  # api_key is optional if CLOUDTRUTH_API_KEY is set
}
```

### 3. Create Resources

```hcl
# Create a project
resource "cloudtruth_project" "app" {
  name        = "my-application"
  description = "Application configuration"
}

# Create a parameter
resource "cloudtruth_parameter" "database_url" {
  name    = "DATABASE_URL"
  project = cloudtruth_project.app.name
  secret  = true
}

# Set environment-specific values
resource "cloudtruth_parameter_value" "db_url_dev" {
  parameter_name = cloudtruth_parameter.database_url.name
  project        = cloudtruth_project.app.name
  environment    = "development"
  value          = "postgresql://localhost:5432/myapp_dev"
}

resource "cloudtruth_parameter_value" "db_url_prod" {
  parameter_name = cloudtruth_parameter.database_url.name
  project        = cloudtruth_project.app.name
  environment    = "production"
  value          = "postgresql://prod.example.com:5432/myapp"
}
```

## Documentation

- **[Provider Documentation](https://registry.terraform.io/providers/cloudtruth/cloudtruth/latest/docs)** - Complete reference for all resources and data sources
- **[CloudTruth Documentation](https://docs.cloudtruth.com/)** - Learn about CloudTruth concepts and features
- **[Examples](./examples/)** - Example configurations for common use cases

### Available Resources

- `cloudtruth_project` - Configuration projects with inheritance
- `cloudtruth_environment` - Deployment environments
- `cloudtruth_parameter` - Configuration parameters
- `cloudtruth_parameter_value` - Environment-specific parameter values
- `cloudtruth_template` - Dynamic configuration templates
- `cloudtruth_type` - Custom parameter types with validation
- `cloudtruth_tag` - Environment tags for versioning
- `cloudtruth_group` - User groups
- `cloudtruth_access_grant` - Access control grants
- `cloudtruth_aws_integration` - AWS integrations
- `cloudtruth_aws_import_action` - AWS import actions
- `cloudtruth_aws_push_action` - AWS push actions
- `cloudtruth_azure_integration` - Azure Key Vault integrations
- `cloudtruth_azure_import_action` - Azure import actions
- `cloudtruth_azure_push_action` - Azure push actions

### Available Data Sources

- `cloudtruth_parameter_value` - Read parameter values
- `cloudtruth_parameter_values` - Read multiple parameter values
- `cloudtruth_template` - Read template content
- `cloudtruth_templates` - List templates
- `cloudtruth_tag` - Read tag information
- `cloudtruth_user` - Read user information
- `cloudtruth_users` - List users

## Development

### Building the Provider

```bash
make build
```

### Running Tests

```bash
# Unit tests
make test

# Acceptance tests (requires CloudTruth credentials)
export CLOUDTRUTH_API_KEY=your-test-api-key
export CLOUDTRUTH_PROJECT=AcceptanceTest
make testacc
```

### Code Quality

```bash
make fmt       # Format code
make fmtcheck  # Check formatting
make lint      # Run linters
make vet       # Run go vet
```

### Generating Documentation

```bash
go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```

### Regenerating API Client

The OpenAPI client is auto-generated. To regenerate from the latest API spec:

```bash
make client
```

## Contributing

We welcome contributions! Please see our contributing guidelines for details on:

- Reporting bugs
- Suggesting enhancements
- Submitting pull requests

### Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and linters (`make test`, `make lint`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to your branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

### Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR-USERNAME/terraform-provider-cloudtruth.git
cd terraform-provider-cloudtruth

# Install dependencies
go mod download

# Run tests
make test
```

## Support

- **Documentation**: [docs.cloudtruth.com](https://docs.cloudtruth.com/)
- **Issues**: [GitHub Issues](https://github.com/cloudtruth/terraform-provider-cloudtruth/issues)
- **Email**: support@cloudtruth.com

## Security

For security concerns, please email security@cloudtruth.com instead of using the public issue tracker.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## About CloudTruth

[CloudTruth](https://cloudtruth.com/) is a universal configuration and secrets management platform that helps teams:

- Centralize configuration across all environments
- Eliminate configuration drift and errors
- Integrate with existing tools and workflows
- Maintain audit trails and access controls
- Sync secrets with AWS, Azure, and other platforms

Learn more at [cloudtruth.com](https://cloudtruth.com/) or try it free at [app.cloudtruth.com](https://app.cloudtruth.com/).

---

**Made with ❤️ by [CloudTruth](https://cloudtruth.com/)**
