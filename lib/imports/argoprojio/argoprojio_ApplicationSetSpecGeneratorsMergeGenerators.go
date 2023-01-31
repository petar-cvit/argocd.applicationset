// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGenerators struct {
	ClusterDecisionResource *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResource `field:"optional" json:"clusterDecisionResource" yaml:"clusterDecisionResource"`
	Clusters *ApplicationSetSpecGeneratorsMergeGeneratorsClusters `field:"optional" json:"clusters" yaml:"clusters"`
	Git *ApplicationSetSpecGeneratorsMergeGeneratorsGit `field:"optional" json:"git" yaml:"git"`
	List *ApplicationSetSpecGeneratorsMergeGeneratorsList `field:"optional" json:"list" yaml:"list"`
	Matrix interface{} `field:"optional" json:"matrix" yaml:"matrix"`
	Merge interface{} `field:"optional" json:"merge" yaml:"merge"`
	PullRequest *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	ScmProvider *ApplicationSetSpecGeneratorsMergeGeneratorsScmProvider `field:"optional" json:"scmProvider" yaml:"scmProvider"`
	Selector *ApplicationSetSpecGeneratorsMergeGeneratorsSelector `field:"optional" json:"selector" yaml:"selector"`
}

