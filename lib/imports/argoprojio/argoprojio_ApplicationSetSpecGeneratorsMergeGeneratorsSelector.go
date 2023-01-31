// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsSelector struct {
	MatchExpressions *[]*ApplicationSetSpecGeneratorsMergeGeneratorsSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

