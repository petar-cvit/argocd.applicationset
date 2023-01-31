// argoprojio
package argoprojio


// OrphanedResourceKey is a reference to a resource to be ignored from.
type AppProjectSpecOrphanedResourcesIgnore struct {
	Group *string `field:"optional" json:"group" yaml:"group"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	Name *string `field:"optional" json:"name" yaml:"name"`
}

