// argoprojio
package argoprojio


// SyncOperationResource contains resources to sync.
type ApplicationOperationSyncResources struct {
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	Name *string `field:"required" json:"name" yaml:"name"`
	Group *string `field:"optional" json:"group" yaml:"group"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

