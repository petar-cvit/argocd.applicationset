// argoprojio
package argoprojio


// Destination is a reference to the target Kubernetes server and namespace.
type ApplicationSpecDestination struct {
	// Name is an alternate way of specifying the target cluster by its symbolic name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace specifies the target namespace for the application's resources.
	//
	// The namespace will only be set for namespace-scoped resources that have not set a value for .metadata.namespace
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Server specifies the URL of the target cluster and must be set to the Kubernetes control plane API.
	Server *string `field:"optional" json:"server" yaml:"server"`
}

