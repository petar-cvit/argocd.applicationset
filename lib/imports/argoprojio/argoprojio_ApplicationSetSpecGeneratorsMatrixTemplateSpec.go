// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixTemplateSpec struct {
	Destination *ApplicationSetSpecGeneratorsMatrixTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecGeneratorsMatrixTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecGeneratorsMatrixTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecGeneratorsMatrixTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecGeneratorsMatrixTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

