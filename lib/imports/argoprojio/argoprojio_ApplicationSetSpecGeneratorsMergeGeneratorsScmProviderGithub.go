// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderGithub struct {
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	AppSecretName *string `field:"optional" json:"appSecretName" yaml:"appSecretName"`
	TokenRef *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderGithubTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

