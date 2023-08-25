package example

//go:generate go run ../cmd/sqlla/main.go

//sqlla:table user_item_detail
//genddl:view user_item_detail
type UserItemDetail struct {
	UserItem UserItem `db:"ui_,nested"`
	Item     Item     `db:"i_,nested"`
	User     User     `db:"u_,nested"`
}

func (u UserItemDetail) _selectStatement() string {
	return `
SELECT
  ui.id, ui.user_id, ui.item_id, ui.is_used, ui.has_extension, ui.used_at,
  i.id, i.name,
  u.id, u.name, u.age, u.rate, u.icon_image, u.created_at, u.updated_at
FROM user_item AS ui
  INNER JOIN item AS i ON ui.item_id = i.id
  INNER JOIN user AS u ON ui.user_id = u.id
`
}
