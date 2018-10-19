package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jimmiebtlr/terraform-provider-kubernetes/kubernetes"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kubernetes.Provider})
}
