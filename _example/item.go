package example

//go:generate go run ../cmd/sqlla/main.go

type ItemID string

//sqlla:table item
//genddl:table item
type Item struct {
	ID   ItemID `db:"id"`
	Name string `db:"name"`
}
