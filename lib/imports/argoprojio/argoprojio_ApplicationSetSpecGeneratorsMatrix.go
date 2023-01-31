// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrix struct {
	Generators *[]*ApplicationSetSpecGeneratorsMatrixGenerators `field:"required" json:"generators" yaml:"generators"`
	Template *ApplicationSetSpecGeneratorsMatrixTemplate `field:"optional" json:"template" yaml:"template"`
}

