// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	BasicAuth *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

