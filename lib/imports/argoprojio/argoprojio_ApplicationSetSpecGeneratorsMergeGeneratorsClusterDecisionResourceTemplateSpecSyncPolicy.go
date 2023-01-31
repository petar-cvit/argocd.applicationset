// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

