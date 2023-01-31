// argoprojio
package argoprojio


// ApplicationSpec represents desired application state.
//
// Contains link to repository with application definition and additional parameters link definition revision.
type ApplicationSpec struct {
	// Destination is a reference to the target Kubernetes server and namespace.
	Destination *ApplicationSpecDestination `field:"required" json:"destination" yaml:"destination"`
	// Project is a reference to the project this application belongs to.
	//
	// The empty string means that application belongs to the 'default' project.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Source is a reference to the location of the application's manifests or chart.
	Source *ApplicationSpecSource `field:"required" json:"source" yaml:"source"`
	// IgnoreDifferences is a list of resources and their fields which should be ignored during comparison.
	IgnoreDifferences *[]*ApplicationSpecIgnoreDifferences `field:"optional" json:"ignoreDifferences" yaml:"ignoreDifferences"`
	// Info contains a list of information (URLs, email addresses, and plain text) that relates to the application.
	Info *[]*ApplicationSpecInfo `field:"optional" json:"info" yaml:"info"`
	// RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions.
	//
	// This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.
	RevisionHistoryLimit *float64 `field:"optional" json:"revisionHistoryLimit" yaml:"revisionHistoryLimit"`
	// SyncPolicy controls when and how a sync will be performed.
	SyncPolicy *ApplicationSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

