package example

import "time"

//go:generate go run ../cmd/sqlla/main.go

// +table: user_sns
//
//sqlla:table user_sns
type UserSNS struct {
	ID        uint64    `db:"id,primarykey"`
	SNSType   string    `db:"sns_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
