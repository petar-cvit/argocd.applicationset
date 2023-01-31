// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

