// argoprojio
package argoprojio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// AppProject provides a logical grouping of applications, providing controls for: * where the apps may deploy to (cluster whitelist) * what may be deployed (repository whitelist, resource whitelist/blacklist) * who can access these applications (roles, OIDC group claims bindings) * and what they can do (RBAC policies) * automation access to these roles (JWT tokens).
type AppProjectProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// AppProjectSpec is the specification of an AppProject.
	Spec *AppProjectSpec `field:"required" json:"spec" yaml:"spec"`
}

