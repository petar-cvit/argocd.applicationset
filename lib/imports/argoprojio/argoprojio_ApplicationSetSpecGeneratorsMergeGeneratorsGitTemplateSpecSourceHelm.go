// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecSourceHelm struct {
	FileParameters *[]*ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecSourceHelmFileParameters `field:"optional" json:"fileParameters" yaml:"fileParameters"`
	IgnoreMissingValueFiles *bool `field:"optional" json:"ignoreMissingValueFiles" yaml:"ignoreMissingValueFiles"`
	Parameters *[]*ApplicationSetSpecGeneratorsMergeGeneratorsGitTemplateSpecSourceHelmParameters `field:"optional" json:"parameters" yaml:"parameters"`
	PassCredentials *bool `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	ReleaseName *string `field:"optional" json:"releaseName" yaml:"releaseName"`
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	ValueFiles *[]*string `field:"optional" json:"valueFiles" yaml:"valueFiles"`
	Values *string `field:"optional" json:"values" yaml:"values"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

