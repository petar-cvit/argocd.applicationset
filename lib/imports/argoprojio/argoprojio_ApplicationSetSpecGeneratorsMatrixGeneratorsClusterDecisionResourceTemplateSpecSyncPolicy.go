// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResourceTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

