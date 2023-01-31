// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeTemplateSpecIgnoreDifferences struct {
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	Group *string `field:"optional" json:"group" yaml:"group"`
	JqPathExpressions *[]*string `field:"optional" json:"jqPathExpressions" yaml:"jqPathExpressions"`
	JsonPointers *[]*string `field:"optional" json:"jsonPointers" yaml:"jsonPointers"`
	ManagedFieldsManagers *[]*string `field:"optional" json:"managedFieldsManagers" yaml:"managedFieldsManagers"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

