package gitrepositorybranch

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the azuredevops_git_repository_branch resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository_branch", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "gitrepositorybranch"

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["repository_id"] = config.Reference{
			Type:      "github.com/chelala/provider-azuredevops/apis/gitrepository/v1alpha1.Repository",
			Extractor: "github.com/chelala/provider-azuredevops/apis/gitrepository/v1alpha1.ExtractResourceID()",
		}
		r.References["ref_branch"] = config.Reference{
			Type: "github.com/chelala/provider-azuredevops/apis/gitrepositorybranch/v1alpha1.RepositoryBranch",
		}
	})
}
