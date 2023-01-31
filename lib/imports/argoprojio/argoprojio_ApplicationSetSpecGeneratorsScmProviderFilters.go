// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderFilters struct {
	BranchMatch *string `field:"optional" json:"branchMatch" yaml:"branchMatch"`
	LabelMatch *string `field:"optional" json:"labelMatch" yaml:"labelMatch"`
	PathsDoNotExist *[]*string `field:"optional" json:"pathsDoNotExist" yaml:"pathsDoNotExist"`
	PathsExist *[]*string `field:"optional" json:"pathsExist" yaml:"pathsExist"`
	RepositoryMatch *string `field:"optional" json:"repositoryMatch" yaml:"repositoryMatch"`
}

