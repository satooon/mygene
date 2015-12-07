package db

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(datasource string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", datasource+"?interpolateParams=true&collation=utf8mb4_bin")
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	return dbMap, nil
}
