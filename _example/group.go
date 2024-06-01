package example

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

//go:generate go run ../cmd/sqlla/main.go

type GroupID uint64

//sqlla:table group
//genddl:table group
type Group struct {
	ID              GroupID          `db:"id,primarykey,autoincrement"`
	Name            string           `db:"name"`
	LeaderUserID    UserId           `db:"leader_user_id"`
	SubLeaderUserID sql.Null[UserId] `db:"sub_leader_user_id"`

	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}