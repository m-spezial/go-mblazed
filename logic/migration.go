package logic

// MigrationManager is an Interface to handle Database Migration.
type MigrationManager interface {
	// GetDBVersions returns all supported and existing Versions.
	GetDBVersions() []string
	// PreMigrate are actions that preform changes to the database schema while different versions of the application
	// are running. This is useful, if you have a high available application without downtime. This function will be
	// called after starting the application.
	PreMigrate(version string) error
	// PostMigrate make changes to the database schema after old versions of the application are stopped. So breaking
	// changes are possible, like deleting tables or columns. This colum need to be called manually through the cli
	PostMigrate(version string) error
}
