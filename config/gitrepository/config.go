package gitrepository

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the azuredevops_git_repository resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {

		r.ShortGroup = "gitrepository"

		r.References["project_id"] = config.Reference{
			Type: "github.com/chelala/provider-azuredevops/apis/project/v1alpha1.Project",
		}
	})
}
