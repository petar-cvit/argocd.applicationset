// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProvider struct {
	AzureDevOps *ApplicationSetSpecGeneratorsScmProviderAzureDevOps `field:"optional" json:"azureDevOps" yaml:"azureDevOps"`
	Bitbucket *ApplicationSetSpecGeneratorsScmProviderBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	BitbucketServer *ApplicationSetSpecGeneratorsScmProviderBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	CloneProtocol *string `field:"optional" json:"cloneProtocol" yaml:"cloneProtocol"`
	Filters *[]*ApplicationSetSpecGeneratorsScmProviderFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsScmProviderGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsScmProviderGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsScmProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsScmProviderTemplate `field:"optional" json:"template" yaml:"template"`
}

