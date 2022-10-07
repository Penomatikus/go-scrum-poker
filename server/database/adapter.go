package database

// Adapter defines an implementation to migrate a schema to a database
type Adapter interface {
	migrateSchema() error
}
