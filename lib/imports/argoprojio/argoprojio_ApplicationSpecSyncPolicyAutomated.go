// argoprojio
package argoprojio


// Automated will keep an application synced to the target revision.
type ApplicationSpecSyncPolicyAutomated struct {
	// AllowEmpty allows apps have zero live resources (default: false).
	AllowEmpty *bool `field:"optional" json:"allowEmpty" yaml:"allowEmpty"`
	// Prune specifies whether to delete resources from the cluster that are not found in the sources anymore as part of automated sync (default: false).
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// SelfHeal specifes whether to revert resources back to their desired state upon modification in the cluster (default: false).
	SelfHeal *bool `field:"optional" json:"selfHeal" yaml:"selfHeal"`
}

