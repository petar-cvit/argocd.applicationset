// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsPullRequestTemplateSpecSource struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourceDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourceHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourceKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourcePlugin `field:"optional" json:"plugin" yaml:"plugin"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

