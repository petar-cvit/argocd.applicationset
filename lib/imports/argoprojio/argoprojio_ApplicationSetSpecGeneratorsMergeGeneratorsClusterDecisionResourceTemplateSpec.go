// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

