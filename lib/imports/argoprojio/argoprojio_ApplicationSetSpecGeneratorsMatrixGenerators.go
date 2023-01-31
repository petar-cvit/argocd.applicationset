// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGenerators struct {
	ClusterDecisionResource *ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResource `field:"optional" json:"clusterDecisionResource" yaml:"clusterDecisionResource"`
	Clusters *ApplicationSetSpecGeneratorsMatrixGeneratorsClusters `field:"optional" json:"clusters" yaml:"clusters"`
	Git *ApplicationSetSpecGeneratorsMatrixGeneratorsGit `field:"optional" json:"git" yaml:"git"`
	List *ApplicationSetSpecGeneratorsMatrixGeneratorsList `field:"optional" json:"list" yaml:"list"`
	Matrix interface{} `field:"optional" json:"matrix" yaml:"matrix"`
	Merge interface{} `field:"optional" json:"merge" yaml:"merge"`
	PullRequest *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	ScmProvider *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProvider `field:"optional" json:"scmProvider" yaml:"scmProvider"`
	Selector *ApplicationSetSpecGeneratorsMatrixGeneratorsSelector `field:"optional" json:"selector" yaml:"selector"`
}

