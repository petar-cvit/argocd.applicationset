// argoprojio
package argoprojio


// Retry controls the strategy to apply if a sync fails.
type ApplicationOperationRetry struct {
	// Backoff controls how to backoff on subsequent retries of failed syncs.
	Backoff *ApplicationOperationRetryBackoff `field:"optional" json:"backoff" yaml:"backoff"`
	// Limit is the maximum number of attempts for retrying a failed sync.
	//
	// If set to 0, no retries will be performed.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

