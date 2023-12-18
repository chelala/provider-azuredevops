package gitrepository

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "gitrepository"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["project_id"] = config.Reference{
			Type: "github.com/chelala/provider-azuredevops/apis/project/v1alpha1.Project",
		}
	})
}
