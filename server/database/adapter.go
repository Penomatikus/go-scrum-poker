package database

import (
	"context"
	"database/sql"
)

// Adapter defines an implementation to migrate a schema to a database
type Adapter interface {
	migrateSchema() error
	mustGetDB(context.Context) *sql.DB
}
