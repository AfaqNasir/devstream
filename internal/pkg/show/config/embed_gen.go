// Code generated by gen_embed_var.go; DO NOT EDIT.
package config

import _ "embed"

//go:embed default.yaml
var DefaultConfig string

// plugin default config
var (

	//go:embed plugins/argocdapp.yaml
	ArgocdappDefaultConfig string

	//go:embed plugins/ci-generic.yaml
	CiGenericDefaultConfig string

	//go:embed plugins/devlake-config.yaml
	DevlakeConfigDefaultConfig string

	//go:embed plugins/github-actions.yaml
	GithubActionsDefaultConfig string

	//go:embed plugins/gitlab-ce-docker.yaml
	GitlabCeDockerDefaultConfig string

	//go:embed plugins/gitlab-ci.yaml
	GitlabCiDefaultConfig string

	//go:embed plugins/harbor-docker.yaml
	HarborDockerDefaultConfig string

	//go:embed plugins/helm-installer.yaml
	HelmInstallerDefaultConfig string

	//go:embed plugins/jenkins-pipeline.yaml
	JenkinsPipelineDefaultConfig string

	//go:embed plugins/jira-github-integ.yaml
	JiraGithubIntegDefaultConfig string

	//go:embed plugins/repo-scaffolding.yaml
	RepoScaffoldingDefaultConfig string

	//go:embed plugins/trello-github-integ.yaml
	TrelloGithubIntegDefaultConfig string

	//go:embed plugins/trello.yaml
	TrelloDefaultConfig string

	//go:embed plugins/zentao.yaml
	ZentaoDefaultConfig string
)

var pluginDefaultConfigs = map[string]string{
	"argocdapp":           ArgocdappDefaultConfig,
	"ci-generic":          CiGenericDefaultConfig,
	"devlake-config":      DevlakeConfigDefaultConfig,
	"github-actions":      GithubActionsDefaultConfig,
	"gitlab-ce-docker":    GitlabCeDockerDefaultConfig,
	"gitlab-ci":           GitlabCiDefaultConfig,
	"harbor-docker":       HarborDockerDefaultConfig,
	"helm-installer":      HelmInstallerDefaultConfig,
	"jenkins-pipeline":    JenkinsPipelineDefaultConfig,
	"jira-github-integ":   JiraGithubIntegDefaultConfig,
	"repo-scaffolding":    RepoScaffoldingDefaultConfig,
	"trello-github-integ": TrelloGithubIntegDefaultConfig,
	"trello":              TrelloDefaultConfig,
	"zentao":              ZentaoDefaultConfig,
}

//go:embed templates/quickstart.yaml
var QuickStart string

//go:embed templates/gitops.yaml
var GitOps string

//go:embed templates/apps.yaml
var Apps string
