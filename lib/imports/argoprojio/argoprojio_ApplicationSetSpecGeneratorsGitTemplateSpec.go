// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsGitTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsGitTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsGitTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsGitTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsGitTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsGitTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

