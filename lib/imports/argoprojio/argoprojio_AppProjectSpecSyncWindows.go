// argoprojio
package argoprojio


// SyncWindow contains the kind, time, duration and attributes that are used to assign the syncWindows to apps.
type AppProjectSpecSyncWindows struct {
	// Applications contains a list of applications that the window will apply to.
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// Clusters contains a list of clusters that the window will apply to.
	Clusters *[]*string `field:"optional" json:"clusters" yaml:"clusters"`
	// Duration is the amount of time the sync window will be open.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Kind defines if the window allows or blocks syncs.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// ManualSync enables manual syncs when they would otherwise be blocked.
	ManualSync *bool `field:"optional" json:"manualSync" yaml:"manualSync"`
	// Namespaces contains a list of namespaces that the window will apply to.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Schedule is the time the window will begin, specified in cron format.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// TimeZone of the sync that will be applied to the schedule.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

