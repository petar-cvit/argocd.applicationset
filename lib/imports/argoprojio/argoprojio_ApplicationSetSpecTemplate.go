// argoprojio
package argoprojio


type ApplicationSetSpecTemplate struct {
	Metadata *ApplicationSetSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	Spec *ApplicationSetSpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

