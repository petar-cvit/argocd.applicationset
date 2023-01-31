// argoprojio
package argoprojio


// SyncPolicy controls when and how a sync will be performed.
type ApplicationSpecSyncPolicy struct {
	// Automated will keep an application synced to the target revision.
	Automated *ApplicationSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	// Retry controls failed sync retry behavior.
	Retry *ApplicationSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	// Options allow you to specify whole app sync-options.
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

