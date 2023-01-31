// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsSelector struct {
	MatchExpressions *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

