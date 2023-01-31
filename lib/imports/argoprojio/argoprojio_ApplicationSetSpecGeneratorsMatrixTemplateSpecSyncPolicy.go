// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMatrixTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

