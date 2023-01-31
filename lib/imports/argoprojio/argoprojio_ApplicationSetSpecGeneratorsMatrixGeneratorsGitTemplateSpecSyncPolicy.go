// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

