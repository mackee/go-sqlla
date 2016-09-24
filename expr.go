package sqlla

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
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

type ExprMultiUint64 struct {
	Column string
	Values []uint64
	Op     Operator
}

func (e ExprMultiUint64) ToSql() (string, []interface{}, error) {
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

type ExprMultiUint32 struct {
	Column string
	Values []uint32
	Op     Operator
}

func (e ExprMultiUint32) ToSql() (string, []interface{}, error) {
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

type ExprMultiInt64 struct {
	Column string
	Values []uint64
	Op     Operator
}

func (e ExprMultiInt64) ToSql() (string, []interface{}, error) {
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

type ExprMultiInt32 struct {
	Column string
	Values []uint32
	Op     Operator
}

func (e ExprMultiInt32) ToSql() (string, []interface{}, error) {
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

type ExprNullInt64 struct {
	Column string
	Value  sql.NullInt64
	Op     Operator
}

func (e ExprNullInt64) ToSql() (string, []interface{}, error) {
	var ops, placeholder string
	var err error
	vs := []interface{}{}
	if !e.Value.Valid {
		ops, err = OpIsNull.ToSql()
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

type ExprMultiNullInt64 struct {
	Column string
	Values []sql.NullInt64
	Op     Operator
}

func (e ExprMultiNullInt64) ToSql() (string, []interface{}, error) {
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

type ExprNullString struct {
	Column string
	Value  sql.NullString
	Op     Operator
}

func (e ExprNullString) ToSql() (string, []interface{}, error) {
	var ops, placeholder string
	var err error
	vs := []interface{}{}
	if !e.Value.Valid {
		ops, err = OpIsNull.ToSql()
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

type ExprMultiNullString struct {
	Column string
	Values []sql.NullString
	Op     Operator
}

func (e ExprMultiNullString) ToSql() (string, []interface{}, error) {
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

type ExprMultiString struct {
	Column string
	Values []string
	Op     Operator
}

func (e ExprMultiString) ToSql() (string, []interface{}, error) {
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

type ExprTime struct {
	Column string
	Value  time.Time
	Op     Operator
}

func (e ExprTime) ToSql() (string, []interface{}, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{e.Value}, nil
}

type ExprMultiTime struct {
	Column string
	Values []time.Time
	Op     Operator
}

func (e ExprMultiTime) ToSql() (string, []interface{}, error) {
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

type ExprNullTime struct {
	Column string
	Value  mysql.NullTime
	Op     Operator
}

func (e ExprNullTime) ToSql() (string, []interface{}, error) {
	var ops, placeholder string
	var err error
	vs := []interface{}{}
	if !e.Value.Valid {
		ops, err = OpIsNull.ToSql()
	} else {
		ops, err = e.Op.ToSql()
		placeholder = " ?"
		vs = append(vs, e.Value)
	}
	if err != nil {
		return "", nil, err
	}

	return e.Column + " " + ops + placeholder, []interface{}{e.Value}, nil
}

type ExprMultiNullTime struct {
	Column string
	Values []mysql.NullTime
	Op     Operator
}

func (e ExprMultiNullTime) ToSql() (string, []interface{}, error) {
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
