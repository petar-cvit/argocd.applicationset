// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsScmProviderAzureDevOps struct {
	AccessTokenRef *ApplicationSetSpecGeneratorsScmProviderAzureDevOpsAccessTokenRef `field:"required" json:"accessTokenRef" yaml:"accessTokenRef"`
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	TeamProject *string `field:"required" json:"teamProject" yaml:"teamProject"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
}

