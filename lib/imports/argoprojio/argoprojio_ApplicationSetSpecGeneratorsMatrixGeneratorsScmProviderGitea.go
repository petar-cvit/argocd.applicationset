// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGitea struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	TokenRef *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGiteaTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

