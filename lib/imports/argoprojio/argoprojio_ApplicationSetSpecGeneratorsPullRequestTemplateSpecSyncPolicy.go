// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequestTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

