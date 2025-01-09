package sqlla

import (
	"fmt"
	"strconv"
)

func NewDialect(dialect string) (Dialect, error) {
	switch dialect {
	case "mysql", "sqlite":
		return MySQLDialect{}, nil
	case "postgresql":
		return PostgreSQLDialect{}, nil
	default:
		return nil, fmt.Errorf("unsupported dialect: %s", dialect)
	}
}

type Dialect interface {
	CQuote() string
	CQuoteBy(column string) string
}

type MySQLDialect struct{}

func (MySQLDialect) CQuote() string {
	return strconv.Quote("`")
}

func (MySQLDialect) CQuoteBy(column string) string {
	return strconv.Quote("`" + column + "`")
}

type PostgreSQLDialect struct{}

func (PostgreSQLDialect) CQuote() string {
	return strconv.Quote(`"`)
}

func (PostgreSQLDialect) CQuoteBy(column string) string {
	return strconv.Quote(`"` + column + `"`)
}
