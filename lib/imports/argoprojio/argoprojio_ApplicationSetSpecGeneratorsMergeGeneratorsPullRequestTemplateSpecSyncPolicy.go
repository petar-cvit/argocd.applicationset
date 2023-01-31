// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

