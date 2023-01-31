// argoprojio
package argoprojio


// SyncStrategy describes how to perform the sync.
type ApplicationOperationSyncSyncStrategy struct {
	// Apply will perform a `kubectl apply` to perform the sync.
	Apply *ApplicationOperationSyncSyncStrategyApply `field:"optional" json:"apply" yaml:"apply"`
	// Hook will submit any referenced resources to perform the sync.
	//
	// This is the default strategy.
	Hook *ApplicationOperationSyncSyncStrategyHook `field:"optional" json:"hook" yaml:"hook"`
}

