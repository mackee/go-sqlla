package example

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

//go:generate go run ../cmd/sqlla/main.go
//go:generate genddl -outpath=./sqlite3.sql -driver=sqlite3

// TeamID is Team's surrogate key
type TeamID int64

// Team is ...
//+table: team
type Team struct {
	ID        TeamID         `db:"id,primarykey,autoincrement"`
	Name      string         `db:"name"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
}
