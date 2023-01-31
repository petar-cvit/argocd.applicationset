// argoprojio
package argoprojio


// Backoff controls how to backoff on subsequent retries of failed syncs.
type ApplicationSpecSyncPolicyRetryBackoff struct {
	// Duration is the amount to back off.
	//
	// Default unit is seconds, but could also be a duration (e.g. "2m", "1h")
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Factor is a factor to multiply the base duration after each failed retry.
	Factor *float64 `field:"optional" json:"factor" yaml:"factor"`
	// MaxDuration is the maximum amount of time allowed for the backoff strategy.
	MaxDuration *string `field:"optional" json:"maxDuration" yaml:"maxDuration"`
}

