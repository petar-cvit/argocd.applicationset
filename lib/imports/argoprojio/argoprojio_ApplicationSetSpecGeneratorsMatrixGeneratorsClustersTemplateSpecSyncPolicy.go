// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsClustersTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsClustersTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsClustersTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

