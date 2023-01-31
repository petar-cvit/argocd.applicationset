// argoprojio
package argoprojio


// Kustomize holds kustomize specific options.
type ApplicationSpecSourceKustomize struct {
	// CommonAnnotations is a list of additional annotations to add to rendered manifests.
	CommonAnnotations *map[string]*string `field:"optional" json:"commonAnnotations" yaml:"commonAnnotations"`
	// CommonLabels is a list of additional labels to add to rendered manifests.
	CommonLabels *map[string]*string `field:"optional" json:"commonLabels" yaml:"commonLabels"`
	// ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps.
	ForceCommonAnnotations *bool `field:"optional" json:"forceCommonAnnotations" yaml:"forceCommonAnnotations"`
	// ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps.
	ForceCommonLabels *bool `field:"optional" json:"forceCommonLabels" yaml:"forceCommonLabels"`
	// Images is a list of Kustomize image override specifications.
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// NamePrefix is a prefix appended to resources for Kustomize apps.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// NameSuffix is a suffix appended to resources for Kustomize apps.
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	// Version controls which version of Kustomize to use for rendering manifests.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

