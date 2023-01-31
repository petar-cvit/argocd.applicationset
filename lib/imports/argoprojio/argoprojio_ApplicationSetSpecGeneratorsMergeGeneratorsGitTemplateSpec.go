// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

