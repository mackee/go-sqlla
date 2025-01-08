package example

import (
	"database/sql"
	"time"
)

//go:generate go run github.com/mackee/go-sqlla/v2/cmd/sqlla --plugins plugins/*.tmpl

type GroupID uint64

//sqlla:table group
//genddl:table group
//sqlla:plugin myrelations key=LeaderUserID:User.ID method=Leader
//sqlla:plugin slice
//sqlla:plugin table get=ID&LeaderUserID list=LeaderUserID&SubLeaderUserID create=Name,LeaderUserID,SubLeaderUserID,ChildGroupID,CreatedAt
type Group struct {
	ID              GroupID         `db:"id,primarykey,autoincrement"`
	Name            string          `db:"name"`
	LeaderUserID    UserId          `db:"leader_user_id"`
	SubLeaderUserID sql.Null[int64] `db:"sub_leader_user_id"`
	ChildGroupID    sql.Null[int64] `db:"child_group_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
