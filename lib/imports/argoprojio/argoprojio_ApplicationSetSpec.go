// argoprojio
package argoprojio


type ApplicationSetSpec struct {
	Generators *[]*ApplicationSetSpecGenerators `field:"required" json:"generators" yaml:"generators"`
	Template *ApplicationSetSpecTemplate `field:"required" json:"template" yaml:"template"`
	GoTemplate *bool `field:"optional" json:"goTemplate" yaml:"goTemplate"`
	SyncPolicy *ApplicationSetSpecSyncPolicy `field:"optional" json:"syncPolicy" yaml:"syncPolicy"`
}

