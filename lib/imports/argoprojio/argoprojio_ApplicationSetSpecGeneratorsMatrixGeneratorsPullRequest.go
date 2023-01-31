// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequest struct {
	BitbucketServer *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	Filters *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplate `field:"optional" json:"template" yaml:"template"`
}

