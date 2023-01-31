// argoprojio
package argoprojio


// Plugin holds config management plugin specific options.
type ApplicationSpecSourcePlugin struct {
	// Env is a list of environment variable entries.
	Env *[]*ApplicationSpecSourcePluginEnv `field:"optional" json:"env" yaml:"env"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

