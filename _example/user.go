package example

//go:generate go run ../cmd/sqlla/main.go

//+table: user
type User struct {
	Id   uint64 `db:"id" primarykey`
	Name string `db:"name"`
}
