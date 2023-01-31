// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsListTemplateSpecSourceDirectory struct {
	Exclude *string `field:"optional" json:"exclude" yaml:"exclude"`
	Include *string `field:"optional" json:"include" yaml:"include"`
	Jsonnet *ApplicationSetSpecGeneratorsListTemplateSpecSourceDirectoryJsonnet `field:"optional" json:"jsonnet" yaml:"jsonnet"`
	Recurse *bool `field:"optional" json:"recurse" yaml:"recurse"`
}

