// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderTemplateSpecSourceDirectory struct {
	Exclude *string `field:"optional" json:"exclude" yaml:"exclude"`
	Include *string `field:"optional" json:"include" yaml:"include"`
	Jsonnet *ApplicationSetSpecGeneratorsScmProviderTemplateSpecSourceDirectoryJsonnet `field:"optional" json:"jsonnet" yaml:"jsonnet"`
	Recurse *bool `field:"optional" json:"recurse" yaml:"recurse"`
}

