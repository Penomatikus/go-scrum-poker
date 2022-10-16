package database

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
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
	*db[*sqlx.DB, *gorm.DB]
}

// db is an abstract definition of A.DB and B.DB
type db[A, B any] struct {
	DB any
}

// ProvideORM takes ormType to determine the used ORM with a given dataSourceName
func ProvideORM(ormTyp ORMType, dataSourceName string) (*ORM, error) {
	return provideORM(ormTyp, dataSourceName)
}

// provideORM hides the implementation details of this pkg and takes ormType to determine the used ORM with a given dataSourceName
func provideORM(ormTyp ORMType, dataSourceName string) (o *ORM, err error) {
	o.db = &db[*sqlx.DB, *gorm.DB]{}
	switch ormTyp {
	case SQLX:
		o.DB, err = sqlx.Open("mysql", dataSourceName)
	case Gorm:
		o.DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	default:
		panic("unsupported ORM")
	}

	return
}
