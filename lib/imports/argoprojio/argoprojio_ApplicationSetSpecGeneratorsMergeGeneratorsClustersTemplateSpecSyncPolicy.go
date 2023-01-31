// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsClustersTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

