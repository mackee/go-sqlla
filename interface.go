package sqlla

import (
	"context"
	"database/sql"
)

// DB is interface like *database/sql.DB
type DB interface {
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row

	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)

	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// Scanner is interface like *database/sql.Row
type Scanner interface {
	Scan(...interface{}) error
}
