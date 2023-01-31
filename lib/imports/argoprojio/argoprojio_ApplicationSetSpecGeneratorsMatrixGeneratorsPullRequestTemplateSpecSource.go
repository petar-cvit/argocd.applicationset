// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSource struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSourceDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSourceHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSourceKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestTemplateSpecSourcePlugin `field:"optional" json:"plugin" yaml:"plugin"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

