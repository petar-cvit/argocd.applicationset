// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

