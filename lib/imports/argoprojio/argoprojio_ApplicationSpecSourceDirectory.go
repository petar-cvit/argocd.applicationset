// argoprojio
package argoprojio


// Directory holds path/directory specific options.
type ApplicationSpecSourceDirectory struct {
	// Exclude contains a glob pattern to match paths against that should be explicitly excluded from being used during manifest generation.
	Exclude *string `field:"optional" json:"exclude" yaml:"exclude"`
	// Include contains a glob pattern to match paths against that should be explicitly included during manifest generation.
	Include *string `field:"optional" json:"include" yaml:"include"`
	// Jsonnet holds options specific to Jsonnet.
	Jsonnet *ApplicationSpecSourceDirectoryJsonnet `field:"optional" json:"jsonnet" yaml:"jsonnet"`
	// Recurse specifies whether to scan a directory recursively for manifests.
	Recurse *bool `field:"optional" json:"recurse" yaml:"recurse"`
}

