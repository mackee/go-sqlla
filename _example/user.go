package example

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mackee/go-genddl/index"
)

//go:generate go run github.com/mackee/go-sqlla/v2/cmd/sqlla
//go:generate go run github.com/mackee/go-genddl/cmd/genddl -outpath=./sqlite3.sql -driver=sqlite3

type UserId uint64

// +table: user
//
//sqlla:plugin count
//sqlla:plugin timeHooks create=CreatedAt sameOnUpdate=UpdatedAt
//sqlla:plugin slice columns=Id,Name keyBy=Id groupBy=Name
type User struct {
	Id        UserId         `db:"id,primarykey,autoincrement"`
	Name      string         `db:"name"`
	Age       sql.NullInt64  `db:"age"`
	Rate      float64        `db:"rate,default=0"`
	IconImage []byte         `db:"icon_image"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}

func (s User) _schemaIndex(methods index.Methods) []index.Definition {
	return []index.Definition{
		methods.Unique(s.Name),
	}
}
