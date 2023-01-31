// argoprojio
package argoprojio


type ApplicationSetSpecGenerators struct {
	ClusterDecisionResource *ApplicationSetSpecGeneratorsClusterDecisionResource `field:"optional" json:"clusterDecisionResource" yaml:"clusterDecisionResource"`
	Clusters *ApplicationSetSpecGeneratorsClusters `field:"optional" json:"clusters" yaml:"clusters"`
	Git *ApplicationSetSpecGeneratorsGit `field:"optional" json:"git" yaml:"git"`
	List *ApplicationSetSpecGeneratorsList `field:"optional" json:"list" yaml:"list"`
	Matrix *ApplicationSetSpecGeneratorsMatrix `field:"optional" json:"matrix" yaml:"matrix"`
	Merge *ApplicationSetSpecGeneratorsMerge `field:"optional" json:"merge" yaml:"merge"`
	PullRequest *ApplicationSetSpecGeneratorsPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	ScmProvider *ApplicationSetSpecGeneratorsScmProvider `field:"optional" json:"scmProvider" yaml:"scmProvider"`
	Selector *ApplicationSetSpecGeneratorsSelector `field:"optional" json:"selector" yaml:"selector"`
}

