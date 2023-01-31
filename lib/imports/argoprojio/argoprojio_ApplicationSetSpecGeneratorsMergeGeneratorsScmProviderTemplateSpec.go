// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

