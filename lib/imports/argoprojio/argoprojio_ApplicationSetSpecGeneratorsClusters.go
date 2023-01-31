// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsClusters struct {
	Selector *ApplicationSetSpecGeneratorsClustersSelector `field:"optional" json:"selector" yaml:"selector"`
	Template *ApplicationSetSpecGeneratorsClustersTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

