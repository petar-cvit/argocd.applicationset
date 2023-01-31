// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucket struct {
	AppPasswordRef *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucketAppPasswordRef `field:"required" json:"appPasswordRef" yaml:"appPasswordRef"`
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	User *string `field:"required" json:"user" yaml:"user"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
}

