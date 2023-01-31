// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	BasicAuth *ApplicationSetSpecGeneratorsScmProviderBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

