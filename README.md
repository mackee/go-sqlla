# go-sqlla
Type safe, reflect free, generative SQL Builder

**THIS IS A ALPHA QUALITY RELEASE. API MAY CHANGE WITHOUT NOTICE.**

## INSTALL

```
$ go get github.com/mackee/go-sqlla/cmd/sqlla
```

## SYNOPSIS

**person.go**:
```go
package table

//go:generate sqlla

//+table: person
type Person struct {
	ID uint64 `db:"id"`
	FirstName string `db:"first_name"
	LastName  string `db:"last_name"
}
```

Run generate:
```
$ ls
person.go
$ go generate
$ ls
person.go person_auto.go
```

Same package as the person.go:
```go

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("failed connect database: %s", err)
	}

	q := NewPersonSQL().Select().ID(uint64(1))
	query, args, err := q.ToSql()
	if err != nil {
		log.Fatalf("query build error: %s", err)
	}

	row := db.QueryRow(query, args...)
	var id uint64
	var firstName, lastName string
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatalf("query exec error: %s", err)
	}
	log.Printf("id=%d, first_name=%s, last_name=%s", id, firstName, lastName)
}
```
