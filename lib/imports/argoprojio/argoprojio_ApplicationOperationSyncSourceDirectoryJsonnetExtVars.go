// argoprojio
package argoprojio


// JsonnetVar represents a variable to be passed to jsonnet during manifest generation.
type ApplicationOperationSyncSourceDirectoryJsonnetExtVars struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Value *string `field:"required" json:"value" yaml:"value"`
	Code *bool `field:"optional" json:"code" yaml:"code"`
}

