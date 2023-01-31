// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsClustersSelector struct {
	MatchExpressions *[]*ApplicationSetSpecGeneratorsClustersSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

