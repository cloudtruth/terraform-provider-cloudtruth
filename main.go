package main

import (
	"flag"

	"github.com/cloudtruth/terraform-provider-cloudtruth/cloudtruth"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
// used by goreleaser
// version string = "dev"
// commit  string = ""
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debugMode,
		ProviderAddr: "registry.terraform.io/terraform-provider-cloudtruth/cloudtruth",
		ProviderFunc: cloudtruth.Provider,
	}
	plugin.Serve(opts)
}
