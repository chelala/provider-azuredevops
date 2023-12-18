/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/chelala/provider-azuredevops/config/gitrepository"
	"github.com/chelala/provider-azuredevops/config/gitrepositorybranch"
	"github.com/chelala/provider-azuredevops/config/project"
)

const (
	resourcePrefix = "azuredevops"
	modulePath     = "github.com/chelala/provider-azuredevops"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("chelala.one"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		gitrepository.Configure,
		gitrepositorybranch.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
