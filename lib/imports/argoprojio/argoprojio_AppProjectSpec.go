// argoprojio
package argoprojio


// AppProjectSpec is the specification of an AppProject.
type AppProjectSpec struct {
	// ClusterResourceBlacklist contains list of blacklisted cluster level resources.
	ClusterResourceBlacklist *[]*AppProjectSpecClusterResourceBlacklist `field:"optional" json:"clusterResourceBlacklist" yaml:"clusterResourceBlacklist"`
	// ClusterResourceWhitelist contains list of whitelisted cluster level resources.
	ClusterResourceWhitelist *[]*AppProjectSpecClusterResourceWhitelist `field:"optional" json:"clusterResourceWhitelist" yaml:"clusterResourceWhitelist"`
	// Description contains optional project description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Destinations contains list of destinations available for deployment.
	Destinations *[]*AppProjectSpecDestinations `field:"optional" json:"destinations" yaml:"destinations"`
	// NamespaceResourceBlacklist contains list of blacklisted namespace level resources.
	NamespaceResourceBlacklist *[]*AppProjectSpecNamespaceResourceBlacklist `field:"optional" json:"namespaceResourceBlacklist" yaml:"namespaceResourceBlacklist"`
	// NamespaceResourceWhitelist contains list of whitelisted namespace level resources.
	NamespaceResourceWhitelist *[]*AppProjectSpecNamespaceResourceWhitelist `field:"optional" json:"namespaceResourceWhitelist" yaml:"namespaceResourceWhitelist"`
	// OrphanedResources specifies if controller should monitor orphaned resources of apps in this project.
	OrphanedResources *AppProjectSpecOrphanedResources `field:"optional" json:"orphanedResources" yaml:"orphanedResources"`
	// PermitOnlyProjectScopedClusters determines whether destinations can only reference clusters which are project-scoped.
	PermitOnlyProjectScopedClusters *bool `field:"optional" json:"permitOnlyProjectScopedClusters" yaml:"permitOnlyProjectScopedClusters"`
	// Roles are user defined RBAC roles associated with this project.
	Roles *[]*AppProjectSpecRoles `field:"optional" json:"roles" yaml:"roles"`
	// SignatureKeys contains a list of PGP key IDs that commits in Git must be signed with in order to be allowed for sync.
	SignatureKeys *[]*AppProjectSpecSignatureKeys `field:"optional" json:"signatureKeys" yaml:"signatureKeys"`
	// SourceNamespaces defines the namespaces application resources are allowed to be created in.
	SourceNamespaces *[]*string `field:"optional" json:"sourceNamespaces" yaml:"sourceNamespaces"`
	// SourceRepos contains list of repository URLs which can be used for deployment.
	SourceRepos *[]*string `field:"optional" json:"sourceRepos" yaml:"sourceRepos"`
	// SyncWindows controls when syncs can be run for apps in this project.
	SyncWindows *[]*AppProjectSpecSyncWindows `field:"optional" json:"syncWindows" yaml:"syncWindows"`
}

