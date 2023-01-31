// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsSelector struct {
	MatchExpressions *[]*ApplicationSetSpecGeneratorsSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

