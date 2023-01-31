// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsGitTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsGitTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsGitTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

