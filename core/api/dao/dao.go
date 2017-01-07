package dao

import (
	"github.com/gocraft/dbr"
)

func GetConn() (*dbr.Connection, error) {
	dbinfo := "postgres://postgres:1234@localhost/kalix?sslmode=disable"

	// create a connection (e.g. "postgres", "mysql", or "sqlite3")
	return dbr.Open("postgres", dbinfo, nil)
}
