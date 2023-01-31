// argoprojio
package argoprojio


// GroupKind specifies a Group and a Kind, but does not force a version.
//
// This is useful for identifying concepts during lookup stages without having partially valid types.
type AppProjectSpecNamespaceResourceBlacklist struct {
	Group *string `field:"required" json:"group" yaml:"group"`
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

