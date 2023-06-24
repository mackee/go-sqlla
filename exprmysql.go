//go:build !tinygo.wasm

package sqlla

import (
	"github.com/go-sql-driver/mysql"
)

type ExprMysqlNullTime struct {
	Column string
	Value  mysql.NullTime
	Op     Operator
}

func (e ExprMysqlNullTime) ToSql() (string, []interface{}, error) {
	var ops, placeholder string
	var err error
	vs := []interface{}{}
	if !e.Value.Valid {
		if e.Op == OpNot {
			ops, err = opIsNotNull.ToSql()
		} else {
			ops, err = opIsNull.ToSql()
		}
	} else {
		ops, err = e.Op.ToSql()
		placeholder = " ?"
		vs = append(vs, e.Value)
	}
	if err != nil {
		return "", nil, err
	}

	return e.Column + " " + ops + placeholder, vs, nil
}

type ExprMultiMysqlNullTime struct {
	Column string
	Values []mysql.NullTime
	Op     Operator
}

func (e ExprMultiMysqlNullTime) ToSql() (string, []interface{}, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]interface{}, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, interface{}(v))
	}
	return e.Column + " " + ops, vs, nil
}
