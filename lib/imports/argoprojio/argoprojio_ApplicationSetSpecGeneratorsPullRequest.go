// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequest struct {
	BitbucketServer *ApplicationSetSpecGeneratorsPullRequestBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	Filters *[]*ApplicationSetSpecGeneratorsPullRequestFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsPullRequestGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsPullRequestGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsPullRequestGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsPullRequestTemplate `field:"optional" json:"template" yaml:"template"`
}

