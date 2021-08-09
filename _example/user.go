package example

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mackee/go-genddl/index"
)

//go:generate go run ../cmd/sqlla/main.go
//go:generate genddl -outpath=./sqlite3.sql -driver=sqlite3

// UserID is User's surrogate key
type UserID uint64

// User is ...
//+table: user
type User struct {
	ID        UserID         `db:"id,primarykey,autoincrement"`
	Name      string         `db:"name"`
	Age       sql.NullInt64  `db:"age"`
	Rate      float64        `db:"rate,default=0"`
	TeamID    TeamID         `db:"team_id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}

func (s User) _schemaIndex(methods index.Methods) []index.Definition {
	return []index.Definition{
		methods.Unique(s.Name),
		methods.Complex(s.TeamID),
	}
}
