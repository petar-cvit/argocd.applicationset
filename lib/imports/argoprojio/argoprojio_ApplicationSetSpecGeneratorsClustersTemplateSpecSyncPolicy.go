// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsClustersTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsClustersTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsClustersTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

