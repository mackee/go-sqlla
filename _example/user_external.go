package example

import "time"

//go:generate go run ../cmd/sqlla/main.go

//+table: user_external
type UserExternal struct {
	Id        uint64    `db:"id,primarykey"`
	UserId    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
