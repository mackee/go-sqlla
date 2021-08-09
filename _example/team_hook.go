package example

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

func (u Team) DefaultInsertHook(q teamInsertSQL) (teamInsertSQL, error) {
	now := time.Now()
	return q.ValueCreatedAt(now), nil
}

func (u Team) DefaultUpdateHook(q teamUpdateSQL) (teamUpdateSQL, error) {
	now := time.Now()
	return q.SetUpdatedAt(mysql.NullTime{Time: now, Valid: true}), nil
}
