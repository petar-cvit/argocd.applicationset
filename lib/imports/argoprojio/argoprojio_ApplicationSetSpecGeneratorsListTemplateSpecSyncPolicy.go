// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsListTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsListTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsListTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

