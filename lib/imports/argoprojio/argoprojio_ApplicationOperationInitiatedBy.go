// argoprojio
package argoprojio


// InitiatedBy contains information about who initiated the operations.
type ApplicationOperationInitiatedBy struct {
	// Automated is set to true if operation was initiated automatically by the application controller.
	Automated *bool `field:"optional" json:"automated" yaml:"automated"`
	// Username contains the name of a user who started operation.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

