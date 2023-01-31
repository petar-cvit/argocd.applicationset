// argoprojio
package argoprojio


// OrphanedResources specifies if controller should monitor orphaned resources of apps in this project.
type AppProjectSpecOrphanedResources struct {
	// Ignore contains a list of resources that are to be excluded from orphaned resources monitoring.
	Ignore *[]*AppProjectSpecOrphanedResourcesIgnore `field:"optional" json:"ignore" yaml:"ignore"`
	// Warn indicates if warning condition should be created for apps which have orphaned resources.
	Warn *bool `field:"optional" json:"warn" yaml:"warn"`
}

