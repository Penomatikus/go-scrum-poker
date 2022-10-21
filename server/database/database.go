package database

type database string

const DatabaseCtxKey database = "DB_Key"

type DatabaseDialect string

const (
	DialectMysql    DatabaseDialect = "mysql"
	DialectPostgres DatabaseDialect = "postgres"
	DialectSqlite3  DatabaseDialect = "sqlite3"
)
