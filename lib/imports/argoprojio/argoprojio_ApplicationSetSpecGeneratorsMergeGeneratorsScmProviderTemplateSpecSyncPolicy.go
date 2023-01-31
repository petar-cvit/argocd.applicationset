// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

