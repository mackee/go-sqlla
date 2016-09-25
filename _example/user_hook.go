package example

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

func (u User) DefaultInsertHook(q userInsertSQL) (userInsertSQL, error) {
	now := time.Now()
	return q.ValueCreatedAt(now), nil
}

func (u User) DefaultUpdateHook(q userUpdateSQL) (userUpdateSQL, error) {
	now := time.Now()
	return q.SetUpdatedAt(mysql.NullTime{now, true}), nil
}
