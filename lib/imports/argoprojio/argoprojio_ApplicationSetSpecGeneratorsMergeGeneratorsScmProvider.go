// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProvider struct {
	AzureDevOps *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderAzureDevOps `field:"optional" json:"azureDevOps" yaml:"azureDevOps"`
	Bitbucket *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	BitbucketServer *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	CloneProtocol *string `field:"optional" json:"cloneProtocol" yaml:"cloneProtocol"`
	Filters *[]*ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplate `field:"optional" json:"template" yaml:"template"`
}

