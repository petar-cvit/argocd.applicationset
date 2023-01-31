// argoprojio
package argoprojio


// Operation contains information about a requested or running operation.
type ApplicationOperation struct {
	// Info is a list of informational items for this operation.
	Info *[]*ApplicationOperationInfo `field:"optional" json:"info" yaml:"info"`
	// InitiatedBy contains information about who initiated the operations.
	InitiatedBy *ApplicationOperationInitiatedBy `field:"optional" json:"initiatedBy" yaml:"initiatedBy"`
	// Retry controls the strategy to apply if a sync fails.
	Retry *ApplicationOperationRetry `field:"optional" json:"retry" yaml:"retry"`
	// Sync contains parameters for the operation.
	Sync *ApplicationOperationSync `field:"optional" json:"sync" yaml:"sync"`
}

