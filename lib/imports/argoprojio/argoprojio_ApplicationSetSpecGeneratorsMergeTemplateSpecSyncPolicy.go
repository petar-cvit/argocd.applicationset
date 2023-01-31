// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

