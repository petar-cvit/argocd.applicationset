// argoprojio
package argoprojio


// HelmParameter is a parameter that's passed to helm template during manifest generation.
type ApplicationSpecSourceHelmParameters struct {
	// ForceString determines whether to tell Helm to interpret booleans and numbers as strings.
	ForceString *bool `field:"optional" json:"forceString" yaml:"forceString"`
	// Name is the name of the Helm parameter.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value is the value for the Helm parameter.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

