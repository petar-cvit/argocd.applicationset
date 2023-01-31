// argoprojio
package argoprojio


// Apply will perform a `kubectl apply` to perform the sync.
type ApplicationOperationSyncSyncStrategyApply struct {
	// Force indicates whether or not to supply the --force flag to `kubectl apply`.
	//
	// The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.
	Force *bool `field:"optional" json:"force" yaml:"force"`
}

