// argoprojio
package argoprojio


// Retry controls failed sync retry behavior.
type ApplicationSpecSyncPolicyRetry struct {
	// Backoff controls how to backoff on subsequent retries of failed syncs.
	Backoff *ApplicationSpecSyncPolicyRetryBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	// Limit is the maximum number of attempts for retrying a failed sync.
	//
	// If set to 0, no retries will be performed.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

