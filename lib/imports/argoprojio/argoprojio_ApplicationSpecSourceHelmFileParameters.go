// argoprojio
package argoprojio


// HelmFileParameter is a file parameter that's passed to helm template during manifest generation.
type ApplicationSpecSourceHelmFileParameters struct {
	// Name is the name of the Helm parameter.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path is the path to the file containing the values for the Helm parameter.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

