// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMerge struct {
	Generators *[]*ApplicationSetSpecGeneratorsMergeGenerators `field:"required" json:"generators" yaml:"generators"`
	MergeKeys *[]*string `field:"required" json:"mergeKeys" yaml:"mergeKeys"`
	Template *ApplicationSetSpecGeneratorsMergeTemplate `field:"optional" json:"template" yaml:"template"`
}

