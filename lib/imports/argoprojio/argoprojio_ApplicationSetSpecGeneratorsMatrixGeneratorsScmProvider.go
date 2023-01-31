// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProvider struct {
	AzureDevOps *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderAzureDevOps `field:"optional" json:"azureDevOps" yaml:"azureDevOps"`
	Bitbucket *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	BitbucketServer *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	CloneProtocol *string `field:"optional" json:"cloneProtocol" yaml:"cloneProtocol"`
	Filters *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplate `field:"optional" json:"template" yaml:"template"`
}

