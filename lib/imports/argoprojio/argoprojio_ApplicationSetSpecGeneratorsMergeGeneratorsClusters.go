// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClusters struct {
	Selector *ApplicationSetSpecGeneratorsMergeGeneratorsClustersSelector `field:"optional" json:"selector" yaml:"selector"`
	Template *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

