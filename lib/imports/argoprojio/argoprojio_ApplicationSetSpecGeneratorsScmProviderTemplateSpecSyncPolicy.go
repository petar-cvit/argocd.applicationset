// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsScmProviderTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsScmProviderTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

