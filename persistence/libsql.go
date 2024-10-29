package persistence

import (
	"database/sql"

	"github.com/schlauerlauer/liesl/db"
	_ "github.com/tursodatabase/go-libsql"
)

func NewRepository(databasePath string) (*db.Queries, error) {

	sqlDb, err := sql.Open("libsql", databasePath)
	if err != nil {
		return nil, err
	}

	queries := db.New(sqlDb)

	_ = sqlDb.QueryRow("PRAGMA journal_mode = WAL;").Scan(nil)
	_ = sqlDb.QueryRow("PRAGMA foreign_keys = on;").Scan(nil)

	return queries, nil
}
