// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsGit struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	Directories *[]*ApplicationSetSpecGeneratorsGitDirectories `field:"optional" json:"directories" yaml:"directories"`
	Files *[]*ApplicationSetSpecGeneratorsGitFiles `field:"optional" json:"files" yaml:"files"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsGitTemplate `field:"optional" json:"template" yaml:"template"`
}

