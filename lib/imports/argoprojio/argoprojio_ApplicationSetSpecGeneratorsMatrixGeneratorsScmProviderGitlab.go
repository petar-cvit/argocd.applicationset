// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGitlab struct {
	Group *string `field:"required" json:"group" yaml:"group"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	IncludeSubgroups *bool `field:"optional" json:"includeSubgroups" yaml:"includeSubgroups"`
	TokenRef *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderGitlabTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

