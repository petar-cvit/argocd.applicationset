// argoprojio
package argoprojio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Application is a definition of Application resource.
type ApplicationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// ApplicationSpec represents desired application state.
	//
	// Contains link to repository with application definition and additional parameters link definition revision.
	Spec *ApplicationSpec `field:"required" json:"spec" yaml:"spec"`
	// Operation contains information about a requested or running operation.
	Operation *ApplicationOperation `field:"optional" json:"operation" yaml:"operation"`
}

