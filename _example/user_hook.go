package example

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

func (u User) DefaultUpdateHook(q userUpdateSQL) (userUpdateSQL, error) {
	now := time.Now()
	return q.SetUpdatedAt(mysql.NullTime{Time: now, Valid: true}), nil
}
