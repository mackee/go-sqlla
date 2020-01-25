module github.com/mackee/go-sqlla/_example

go 1.13

replace github.com/mackee/go-sqlla => ../

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/mackee/go-genddl v0.0.0-20181004101258-0e8f5ffc20f4
	github.com/mackee/go-sqlla v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
)
