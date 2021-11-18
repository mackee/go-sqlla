package example

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

func (u User) DefaultInsertHook(q userInsertSQL) (userInsertSQL, error) {
	now := time.Now()
	return q.ValueCreatedAt(now), nil
}

func (u User) DefaultInsertOnDuplicateKeyUpdateHook(q userInsertOnDuplicateKeyUpdateSQL) (userInsertOnDuplicateKeyUpdateSQL, error) {
	now := time.Now()
	q.insertSQL = q.insertSQL.ValueUpdatedAt(mysql.NullTime{Time: now, Valid: true})

	return q.SameOnUpdateUpdatedAt(), nil
}

func (u User) DefaultUpdateHook(q userUpdateSQL) (userUpdateSQL, error) {
	now := time.Now()
	return q.SetUpdatedAt(mysql.NullTime{Time: now, Valid: true}), nil
}
