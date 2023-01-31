// argoprojio
package argoprojio


// EnvEntry represents an entry in the application's environment.
type ApplicationOperationSyncSourcePluginEnv struct {
	// Name is the name of the variable, usually expressed in uppercase.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value is the value of the variable.
	Value *string `field:"required" json:"value" yaml:"value"`
}

