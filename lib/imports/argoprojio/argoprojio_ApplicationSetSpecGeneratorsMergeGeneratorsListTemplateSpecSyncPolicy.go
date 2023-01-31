// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsListTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

