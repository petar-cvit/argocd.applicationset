// argoprojio
package argoprojio


type ApplicationSetSpecTemplateSpec struct {
	Destination *ApplicationSetSpecTemplateSpecDestination `field:"required" json:"destination" yaml:"destination"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Source *ApplicationSetSpecTemplateSpecSource `field:"required" json:"source" yaml:"source"`
	IgnoreDifferences *[]*ApplicationSetSpecTemplateSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	Info *[]*ApplicationSetSpecTemplateSpecInfo `field:"optional" json:"info" yaml:"info"`
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	SyncPolicy *ApplicationSetSpecTemplateSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

