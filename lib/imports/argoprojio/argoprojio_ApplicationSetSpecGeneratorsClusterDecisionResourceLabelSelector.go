// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsClusterDecisionResourceLabelSelector struct {
	MatchExpressions *[]*ApplicationSetSpecGeneratorsClusterDecisionResourceLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

