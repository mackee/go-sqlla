package postgresql

import (
	"database/sql"
	"time"
)

type GroupID int64

//sqlla:table groups
//genddl:table groups
type Group struct {
	ID                 GroupID             `db:"id,primarykey,autoincrement"`
	Name               string              `db:"name"`
	LeaderAccountID    AccountID           `db:"leader_account_id"`
	SubLeaderAccountID sql.Null[AccountID] `db:"sub_leader_account_id"`
	ChildGroupID       sql.Null[GroupID]   `db:"child_group_id"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
