// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourceKustomize struct {
	CommonAnnotations *map[string]*string `field:"optional" json:"commonAnnotations" yaml:"commonAnnotations"`
	CommonLabels *map[string]*string `field:"optional" json:"commonLabels" yaml:"commonLabels"`
	ForceCommonAnnotations *bool `field:"optional" json:"forceCommonAnnotations" yaml:"forceCommonAnnotations"`
	ForceCommonLabels *bool `field:"optional" json:"forceCommonLabels" yaml:"forceCommonLabels"`
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

