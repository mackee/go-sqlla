package example

//go:generate go run ../cmd/sqlla/main.go

//+table: user_item
type UserItem struct {
	Id     uint64 `db:"id,primarykey,autoincrement"`
	UserId uint64 `db:"user_id"`
	ItemId string `db:"item_id"`
}
