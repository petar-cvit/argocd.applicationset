// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResource struct {
	ConfigMapRef *string `field:"required" json:"configMapRef" yaml:"configMapRef"`
	LabelSelector *ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResourceLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResourceTemplate `field:"optional" json:"template" yaml:"template"`
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

