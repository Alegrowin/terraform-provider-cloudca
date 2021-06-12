package main

import (
	cloudca "github.com/cloud-ca/terraform-provider-cloudca/cloudca"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return cloudca.Provider()
		},
	})
}
