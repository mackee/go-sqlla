package sqlla

import (
	"strings"
)

var (
	OpEqual        Operator = "="
	OpGreater      Operator = ">"
	OpGreaterEqual Operator = ">="
	OpLess         Operator = "<"
	OpLessEqual    Operator = "<="
	OpNot          Operator = "<>"
	OpIs           Operator = "IS"
	OpIsNull       Operator = "IS NULL"
	OpIsNotNull    Operator = "IS NOT NULL"
)

type Operator string

func (op Operator) ToSql() (string, error) {
	return string(op), nil
}

func MakeInOperator(n int) Operator {
	return Operator("IN(?" + strings.Repeat(",?", n-1) + ")")
}
