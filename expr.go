package sqlla

import (
	"strconv"
)

type Expr interface {
	ToSql() (string, []interface{}, error)
}

type ExprUint64 struct {
	Column string
	Value  uint64
	Op     Operator
}

func (e ExprUint64) ToSql() (string, []interface{}, error) {
	s := strconv.FormatUint(e.Value, 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprUint32 struct {
	Column string
	Value  uint32
	Op     Operator
}

func (e ExprUint32) ToSql() (string, []interface{}, error) {
	s := strconv.FormatUint(uint64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprInt64 struct {
	Column string
	Value  int64
	Op     Operator
}

func (e ExprInt64) ToSql() (string, []interface{}, error) {
	s := strconv.FormatInt(e.Value, 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprInt32 struct {
	Column string
	Value  int32
	Op     Operator
}

func (e ExprInt32) ToSql() (string, []interface{}, error) {
	s := strconv.FormatInt(int64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprString struct {
	Column string
	Value  string
	Op     Operator
}

func (e ExprString) ToSql() (string, []interface{}, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{e.Value}, nil
}
