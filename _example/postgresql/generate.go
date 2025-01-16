package postgresql

//go:generate go run github.com/mackee/go-sqlla/v2/cmd/sqlla --dialect=postgresql --dir-all
//go:generate go run github.com/mackee/go-genddl/cmd/genddl -outpath=./postgresql.sql -driver=pg
