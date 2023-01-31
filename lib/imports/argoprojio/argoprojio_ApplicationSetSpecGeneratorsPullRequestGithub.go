// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequestGithub struct {
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	AppSecretName *string `field:"optional" json:"appSecretName" yaml:"appSecretName"`
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	TokenRef *ApplicationSetSpecGeneratorsPullRequestGithubTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

