// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixTemplateSpecSourceDirectory struct {
	Exclude *string `field:"optional" json:"exclude" yaml:"exclude"`
	Include *string `field:"optional" json:"include" yaml:"include"`
	Jsonnet *ApplicationSetSpecGeneratorsMatrixTemplateSpecSourceDirectoryJsonnet `field:"optional" json:"jsonnet" yaml:"jsonnet"`
	Recurse *bool `field:"optional" json:"recurse" yaml:"recurse"`
}

