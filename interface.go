package sqlla

import "database/sql"

type DB interface {
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

type Scanner interface {
	Scan(...interface{}) error
}
