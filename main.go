package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mcoffin/terraform-provider-localfile/localfile"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: localfile.Provider,
	})
}
