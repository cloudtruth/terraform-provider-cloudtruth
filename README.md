# CloudTruth Terraform Provider

This is the official Terraform provider for [CloudTruth](https://cloudtruth.com/)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.15.x
-	[Go](https://golang.org/doc/install) >= 1.20

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:
```sh
$ go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

If you run the tests in a fork of this repo, you will need to recreate the boilerplate CloudTruth resources which exist in the dedicated
acceptance test account. We will provide a bootstrapping script to ensure that these resource exist in your account.

Additionally, you will need to specify the target CloudTruth project and environment(s) as well as your CloudTruth API key. You can
specify the API key with a `TF_VAR_cloudtruth_api_key` environment variable or with a `CLOUDTRUTH_API_KEY` environment variable (a la the CloudTruth
CLI). You can specify the the project and/or environments inline in your HCL files or via the `CLOUDTRUTH_PROJECT` and `CLOUDTRUTH_ENVIRONMENT`
environment variables.

```sh
$ make testacc
```

### Commit hooks
To enable the commit hooks stored in `.githooks` in your local repo, run `git config --local core.hooksPath .githooks/`
