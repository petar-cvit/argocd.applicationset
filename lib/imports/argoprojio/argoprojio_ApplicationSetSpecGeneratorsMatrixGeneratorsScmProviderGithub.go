// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGithub struct {
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	AppSecretName *string `field:"optional" json:"appSecretName" yaml:"appSecretName"`
	TokenRef *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGithubTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

