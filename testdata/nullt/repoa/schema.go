package repoa

import (
	"database/sql"
	"time"
)

type NullableTime sql.Null[time.Time]

//sqlla:table product
type Product struct {
	ID         uint64              `db:"id,primarykey,autoincrement"`
	ModifiedAt sql.Null[time.Time] `db:"modified_at"`
	RemovedAt  NullableTime        `db:"removed_at"`
}
