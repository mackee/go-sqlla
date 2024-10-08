package sqlla

import (
	"strings"
)

var (
	OpEqual        Operator = "="  // Operator for equal. Column(value, sqlla.OpEqual) same as Column = value
	OpGreater      Operator = ">"  // Operator for greater. Column(value, sqlla.OpGreater) same as Column > value
	OpGreaterEqual Operator = ">=" // Operator for greater equal. Column(value, sqlla.OpGreaterEqual) same as Column >= value
	OpLess         Operator = "<"  // Operator for less. Column(value, sqlla.OpLess) same as Column < value
	OpLessEqual    Operator = "<=" // Operator for less equal. Column(value, sqlla.OpLessEqual) same as Column <= value
	OpNot          Operator = "<>" // Operator for not equal. Column(value, sqlla.OpNot) same as Column <> value
	OpIs           Operator = "IS" // Operator for is. Column(value, sqlla.OpIs) same as Column IS value
	opIsNull       Operator = "IS NULL"
	opIsNotNull    Operator = "IS NOT NULL"
	OpLike         Operator = "LIKE" // Operator for like. Column(value, sqlla.OpLike) same as Column LIKE value
)

type Operator string

func (op Operator) ToSql() (string, error) {
	return string(op), nil
}

func MakeInOperator(n int) Operator {
	return Operator("IN(?" + strings.Repeat(",?", n-1) + ")")
}
