// argoprojio
package argoprojio


// Hook will submit any referenced resources to perform the sync.
//
// This is the default strategy.
type ApplicationOperationSyncSyncStrategyHook struct {
	// Force indicates whether or not to supply the --force flag to `kubectl apply`.
	//
	// The --force flag deletes and re-create the resource, when PATCH encounters conflict and has retried for 5 times.
	Force *bool `field:"optional" json:"force" yaml:"force"`
}

