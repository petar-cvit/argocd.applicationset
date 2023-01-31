// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsGit struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	Directories *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsGitDirectories `field:"optional" json:"directories" yaml:"directories"`
	Files *[]*ApplicationSetSpecGeneratorsMatrixGeneratorsGitFiles `field:"optional" json:"files" yaml:"files"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMatrixGeneratorsGitTemplate `field:"optional" json:"template" yaml:"template"`
}

