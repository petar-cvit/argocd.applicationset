// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestGitlab struct {
	Project *string `field:"required" json:"project" yaml:"project"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	PullRequestState *string `field:"optional" json:"pullRequestState" yaml:"pullRequestState"`
	TokenRef *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestGitlabTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

