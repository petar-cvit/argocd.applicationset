// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	BasicAuth *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

