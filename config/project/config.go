package project

import "github.com/crossplane/upjet/pkg/config"

// Configure configures the azuredevops_project resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_project", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "project"
	})
}
