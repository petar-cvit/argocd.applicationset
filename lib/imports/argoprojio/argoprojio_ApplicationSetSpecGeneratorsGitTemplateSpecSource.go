// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsGitTemplateSpecSource struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsGitTemplateSpecSourceDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsGitTemplateSpecSourceHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsGitTemplateSpecSourceKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsGitTemplateSpecSourcePlugin `field:"optional" json:"plugin" yaml:"plugin"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

