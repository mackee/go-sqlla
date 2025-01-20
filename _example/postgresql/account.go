package postgresql

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type AccountID int64

//genddl:table accounts
//sqlla:table accounts
type Account struct {
	ID        AccountID       `db:"id,autoincrement,primarykey"`
	Name      string          `db:"name"`
	Embedding pgvector.Vector `db:"embedding,type=vector(3)"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}

var FixedNow time.Time

func (a Account) DefaultInsertHook(q accountInsertSQL) (accountInsertSQL, error) {
	now := FixedNow
	if FixedNow.IsZero() {
		now = time.Now()
	}
	return q.
		ValueCreatedAt(now).
		ValueUpdatedAt(now), nil
}

func (a Account) DefaultUpdateHook(q accountUpdateSQL) (accountUpdateSQL, error) {
	now := FixedNow
	if FixedNow.IsZero() {
		now = time.Now()
	}
	return q.
		SetUpdatedAt(now), nil
}
