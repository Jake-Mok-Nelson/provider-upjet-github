package repositorycollaborator

import "github.com/crossplane/upjet/pkg/config"

// Configure github_repository_collaborator resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_repository_collaborator", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.Kind = "RepositoryCollaborator"
		r.ShortGroup = "repo"

		r.References["repository"] = config.Reference{
			TerraformName: "github_repository",
		}
	})
}
