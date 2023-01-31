// argoprojio
package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSource struct {
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	Directory *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSourceDirectory `field:"optional" json:"directory" yaml:"directory"`
	Helm *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSourceHelm `field:"optional" json:"helm" yaml:"helm"`
	Kustomize *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSourceKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	Path *string `field:"optional" json:"path" yaml:"path"`
	Plugin *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplateSpecSourcePlugin `field:"optional" json:"plugin" yaml:"plugin"`
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

