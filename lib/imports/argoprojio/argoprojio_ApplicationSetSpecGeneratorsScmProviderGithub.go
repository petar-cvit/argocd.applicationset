// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderGithub struct {
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	AppSecretName *string `field:"optional" json:"appSecretName" yaml:"appSecretName"`
	TokenRef *ApplicationSetSpecGeneratorsScmProviderGithubTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

