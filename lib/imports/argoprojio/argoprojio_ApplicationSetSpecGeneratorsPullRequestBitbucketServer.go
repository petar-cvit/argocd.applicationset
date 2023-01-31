// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequestBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	BasicAuth *ApplicationSetSpecGeneratorsPullRequestBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

