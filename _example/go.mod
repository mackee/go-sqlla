module github.com/mackee/go-sqlla/v2/_example

go 1.13

replace github.com/mackee/go-sqlla/v2 => ../

require (
	github.com/go-sql-driver/mysql v1.4.0
	github.com/mackee/go-genddl v0.0.0-20181004101258-0e8f5ffc20f4
	github.com/mackee/go-sqlla/v2 v2.5.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
)
