package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database string

const DatabaseCtxKey database = "DB_Key"

func ProvideDatabase() *gorm.DB {
	//TODO: make this configurable
	dataSourceName := "scrumpoker:password@tcp(localhost:3306)/database?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	return db
}

func MustGetDB(ctx context.Context) *gorm.DB {
	if db := ctx.Value(DatabaseCtxKey).(*gorm.DB); db != nil {
		return db
	}
	panic(fmt.Errorf("no DB instance was found during transaktion"))
}
