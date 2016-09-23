package example

import "database/sql"

//go:generate go run ../cmd/sqlla/main.go
//go:generate genddl -outpath=./sqlite3.sql -driver=sqlite3

//+table: user
type User struct {
	Id   uint64        `db:"id,primarykey,autoincrement"`
	Name string        `db:"name"`
	Age  sql.NullInt64 `db:"age"`
}
