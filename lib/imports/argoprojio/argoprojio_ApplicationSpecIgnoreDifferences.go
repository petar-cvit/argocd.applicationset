// argoprojio
package argoprojio


// ResourceIgnoreDifferences contains resource filter and list of json paths which should be ignored during comparison with live state.
type ApplicationSpecIgnoreDifferences struct {
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	Group *string `field:"optional" json:"group" yaml:"group"`
	JqPathExpressions *[]*string `field:"optional" json:"jqPathExpressions" yaml:"jqPathExpressions"`
	JsonPointers *[]*string `field:"optional" json:"jsonPointers" yaml:"jsonPointers"`
	// ManagedFieldsManagers is a list of trusted managers.
	//
	// Fields mutated by those managers will take precedence over the desired state defined in the SCM and won't be displayed in diffs.
	ManagedFieldsManagers *[]*string `field:"optional" json:"managedFieldsManagers" yaml:"managedFieldsManagers"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

