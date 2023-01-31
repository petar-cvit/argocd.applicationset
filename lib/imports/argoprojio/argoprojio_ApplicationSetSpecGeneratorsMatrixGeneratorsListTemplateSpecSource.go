// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSource struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourceDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourceHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourceKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSourcePlugin `field:"optional" json:"plugin" yaml:"plugin"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

