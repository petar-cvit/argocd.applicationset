// argoprojio
package argoprojio


// Jsonnet holds options specific to Jsonnet.
type ApplicationSpecSourceDirectoryJsonnet struct {
	// ExtVars is a list of Jsonnet External Variables.
	ExtVars *[]*ApplicationSpecSourceDirectoryJsonnetExtVars `field:"optional" json:"extVars" yaml:"extVars"`
	// Additional library search dirs.
	Libs *[]*string `field:"optional" json:"libs" yaml:"libs"`
	// TLAS is a list of Jsonnet Top-level Arguments.
	Tlas *[]*ApplicationSpecSourceDirectoryJsonnetTlas `field:"optional" json:"tlas" yaml:"tlas"`
}

