package sqlla

import (
	"bytes"
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
	Values []int32
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

type ExprNullFloat64 struct {
	Column string
	Value  sql.NullFloat64
	Op     Operator
}

func (e ExprNullFloat64) ToSql() (string, []interface{}, error) {
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

type ExprFloat64 struct {
	Column string
	Value  float64
	Op     Operator
}

func (e ExprFloat64) ToSql() (string, []interface{}, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{e.Value}, nil
}

type ExprMultiFloat64 struct {
	Column string
	Values []float64
	Op     Operator
}

func (e ExprMultiFloat64) ToSql() (string, []interface{}, error) {
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

type ExprMultiNullFloat64 struct {
	Column string
	Values []sql.NullFloat64
	Op     Operator
}

func (e ExprMultiNullFloat64) ToSql() (string, []interface{}, error) {
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

type ExprNullBool struct {
	Column string
	Value  sql.NullBool
	Op     Operator
}

func (e ExprNullBool) ToSql() (string, []interface{}, error) {
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

type ExprBool struct {
	Column string
	Value  bool
	Op     Operator
}

func (e ExprBool) ToSql() (string, []interface{}, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{e.Value}, nil
}

type ExprMultiBool struct {
	Column string
	Values []bool
	Op     Operator
}

func (e ExprMultiBool) ToSql() (string, []interface{}, error) {
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

type ExprMultiNullBool struct {
	Column string
	Values []sql.NullBool
	Op     Operator
}

func (e ExprMultiNullBool) ToSql() (string, []interface{}, error) {
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

type ExprOr []Where

func (e ExprOr) ToSql() (string, []interface{}, error) {
	if len(e) == 0 {
		return "", []interface{}{}, nil
	}

	b := new(bytes.Buffer)
	vs := make([]interface{}, 0, 10)
	b.WriteString("(")
	for i, w := range e {
		if i > 0 {
			b.WriteString(" OR ")
		}
		b.WriteString("(")
		q, wvs, err := w.ToSql()
		if err != nil {
			return "", nil, err
		}
		b.WriteString(q)
		b.WriteString(" )")
		vs = append(vs, wvs...)
	}

	b.WriteString(")")

	return b.String(), vs, nil
}

type ExprUint8 struct {
	Column string
	Value  uint8
	Op     Operator
}

func (e ExprUint8) ToSql() (string, []interface{}, error) {
	s := strconv.FormatUint(uint64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprMultiUint8 struct {
	Column string
	Values []uint8
	Op     Operator
}

func (e ExprMultiUint8) ToSql() (string, []interface{}, error) {
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

type ExprInt8 struct {
	Column string
	Value  int8
	Op     Operator
}

func (e ExprInt8) ToSql() (string, []interface{}, error) {
	s := strconv.FormatInt(int64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []interface{}{s}, nil
}

type ExprMultiInt8 struct {
	Column string
	Values []int8
	Op     Operator
}

func (e ExprMultiInt8) ToSql() (string, []interface{}, error) {
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
