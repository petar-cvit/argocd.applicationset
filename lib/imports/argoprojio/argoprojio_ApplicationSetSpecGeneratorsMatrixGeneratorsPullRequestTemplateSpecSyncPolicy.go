// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

