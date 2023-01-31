// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	BasicAuth *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

