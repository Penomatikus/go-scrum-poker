package database

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ORMType provides an enum for all supported go ORMs (or sql interfaces)
type ORMType int

const (
	// For usage of sqlX
	SQLX ORMType = iota
	// For usage of Gorm ORM
	Gorm
)

// ORM defines an adapter for differnet ORM database objects
type ORM struct {
	db[*sqlx.DB, *gorm.DB]
}

// db is an abstract definition of A.DB and B.DB
type db[A, B any] struct {
	DB any
}

// ProvideORM takes ormType to determine the used ORM with a given dataSourceName
func ProvideORM(ormType ORMType, dialect DatabaseDialect, dataSourceName string) (*ORM, error) {
	return provideORM(ormType, dialect, dataSourceName)
}

// provideORM hides the implementation details of this pkg and takes ormType to determine the used ORM with a given dataSourceName
func provideORM(ormType ORMType, dialect DatabaseDialect, dsn string) (o *ORM, err error) {
	o = &ORM{db[*sqlx.DB, *gorm.DB]{}}
	switch ormType {
	case SQLX:
		o.DB, err = sqlx.Open(string(dialect), dsn)
	case Gorm:
		o.DB, err = gorm.Open(provideGormDialector(dialect, dsn), &gorm.Config{})
	default:
		panic("unsupported ORM")
	}

	return
}

func provideGormDialector(dialect DatabaseDialect, dsn string) gorm.Dialector {
	switch dialect {
	case MysqlDialect:
		return mysql.Open(dsn)
	case PostgresDialect:
		return postgres.Open(dsn)
	case Sqlite3Dialect:
		return sqlite.Open(dsn)
	default:
		panic("unsported gorm dialect")
	}
}
