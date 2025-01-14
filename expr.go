package sqlla

import (
	"bytes"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type Expr interface {
	ToSql() (string, []any, error)
}

type ExprWithExprValue interface {
	Expr
	ExprValue() ExprToSqlPger
}

type ExprToSqlPger interface {
	ToSqlPg(int) (string, int, []any, error)
}

type ExprUint64 struct {
	Column string
	Value  uint64
	Op     OperatorToSqler
}

func (e ExprUint64) ToSql() (string, []any, error) {
	s := strconv.FormatUint(e.Value, 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprUint64) ExprValue() ExprValue[uint64] {
	return ExprValue[uint64](e)
}

type ExprMultiUint64 struct {
	Column string
	Values []uint64
	Op     OperatorToSqler
}

func (e ExprMultiUint64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiUint64) ExprValue() ExprMultiValue[uint64] {
	return ExprMultiValue[uint64](e)
}

type ExprUint32 struct {
	Column string
	Value  uint32
	Op     OperatorToSqler
}

func (e ExprUint32) ToSql() (string, []any, error) {
	s := strconv.FormatUint(uint64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprUint32) ExprValue() ExprValue[uint32] {
	return ExprValue[uint32](e)
}

type ExprMultiUint32 struct {
	Column string
	Values []uint32
	Op     OperatorToSqler
}

func (e ExprMultiUint32) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiUint32) ExprValue() ExprMultiValue[uint32] {
	return ExprMultiValue[uint32](e)
}

type ExprInt64 struct {
	Column string
	Value  int64
	Op     OperatorToSqler
}

func (e ExprInt64) ToSql() (string, []any, error) {
	s := strconv.FormatInt(e.Value, 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprInt64) ExprValue() ExprValue[int64] {
	return ExprValue[int64](e)
}

type ExprMultiInt64 struct {
	Column string
	Values []int64
	Op     OperatorToSqler
}

func (e ExprMultiInt64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiInt64) ExprValue() ExprMultiValue[int64] {
	return ExprMultiValue[int64](e)
}

type ExprInt32 struct {
	Column string
	Value  int32
	Op     OperatorToSqler
}

func (e ExprInt32) ToSql() (string, []any, error) {
	s := strconv.FormatInt(int64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprInt32) ExprValue() ExprValue[int32] {
	return ExprValue[int32](e)
}

type ExprMultiInt32 struct {
	Column string
	Values []int32
	Op     OperatorToSqler
}

func (e ExprMultiInt32) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiInt32) ExprValue() ExprMultiValue[int32] {
	return ExprMultiValue[int32](e)
}

type ExprNullInt64 struct {
	Column string
	Value  sql.NullInt64
	Op     OperatorToSqler
}

func (e ExprNullInt64) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNullInt64) ExprValue() ExprValue[sql.NullInt64] {
	return ExprValue[sql.NullInt64](e)
}

type ExprMultiNullInt64 struct {
	Column string
	Values []sql.NullInt64
	Op     OperatorToSqler
}

func (e ExprMultiNullInt64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiNullInt64) ExprValue() ExprMultiValue[sql.NullInt64] {
	return ExprMultiValue[sql.NullInt64](e)
}

type ExprNullString struct {
	Column string
	Value  sql.NullString
	Op     OperatorToSqler
}

func (e ExprNullString) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNullString) ExprValue() ExprValue[sql.NullString] {
	return ExprValue[sql.NullString](e)
}

type ExprMultiNullString struct {
	Column string
	Values []sql.NullString
	Op     OperatorToSqler
}

func (e ExprMultiNullString) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiNullString) ExprValue() ExprMultiValue[sql.NullString] {
	return ExprMultiValue[sql.NullString](e)
}

type ExprString struct {
	Column string
	Value  string
	Op     OperatorToSqler
}

func (e ExprString) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{e.Value}, nil
}

type ExprMultiString struct {
	Column string
	Values []string
	Op     OperatorToSqler
}

func (e ExprMultiString) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiString) ExprValue() ExprMultiValue[string] {
	return ExprMultiValue[string](e)
}

type ExprTime struct {
	Column string
	Value  time.Time
	Op     OperatorToSqler
}

func (e ExprTime) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{e.Value}, nil
}

func (e ExprTime) ExprValue() ExprValue[time.Time] {
	return ExprValue[time.Time](e)
}

type ExprMultiTime struct {
	Column string
	Values []time.Time
	Op     OperatorToSqler
}

func (e ExprMultiTime) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiTime) ExprValue() ExprMultiValue[time.Time] {
	return ExprMultiValue[time.Time](e)
}

type ExprNullTime struct {
	Column string
	Value  sql.NullTime
	Op     OperatorToSqler
}

func (e ExprNullTime) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNullTime) ExprValue() ExprValue[sql.NullTime] {
	return ExprValue[sql.NullTime](e)
}

type ExprMultiNullTime struct {
	Column string
	Values []sql.NullTime
	Op     OperatorToSqler
}

func (e ExprMultiNullTime) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiNullTime) ExprValue() ExprMultiValue[sql.NullTime] {
	return ExprMultiValue[sql.NullTime](e)
}

type ExprNullFloat64 struct {
	Column string
	Value  sql.NullFloat64
	Op     OperatorToSqler
}

func (e ExprNullFloat64) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNullFloat64) ExprValue() ExprValue[sql.NullFloat64] {
	return ExprValue[sql.NullFloat64](e)
}

type ExprFloat64 struct {
	Column string
	Value  float64
	Op     OperatorToSqler
}

func (e ExprFloat64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{e.Value}, nil
}

func (e ExprFloat64) ExprValue() ExprValue[float64] {
	return ExprValue[float64](e)
}

type ExprMultiFloat64 struct {
	Column string
	Values []float64
	Op     OperatorToSqler
}

func (e ExprMultiFloat64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiFloat64) ExprValue() ExprMultiValue[float64] {
	return ExprMultiValue[float64](e)
}

type ExprMultiNullFloat64 struct {
	Column string
	Values []sql.NullFloat64
	Op     OperatorToSqler
}

func (e ExprMultiNullFloat64) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiNullFloat64) ExprValue() ExprMultiValue[sql.NullFloat64] {
	return ExprMultiValue[sql.NullFloat64](e)
}

type ExprNullBool struct {
	Column string
	Value  sql.NullBool
	Op     OperatorToSqler
}

func (e ExprNullBool) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNullBool) ExprValue() ExprValue[sql.NullBool] {
	return ExprValue[sql.NullBool](e)
}

type ExprBool struct {
	Column string
	Value  bool
	Op     OperatorToSqler
}

func (e ExprBool) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{e.Value}, nil
}

func (e ExprBool) ExprValue() ExprValue[bool] {
	return ExprValue[bool](e)
}

type ExprMultiBool struct {
	Column string
	Values []bool
	Op     OperatorToSqler
}

func (e ExprMultiBool) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiBool) ExprValue() ExprMultiValue[bool] {
	return ExprMultiValue[bool](e)
}

type ExprMultiNullBool struct {
	Column string
	Values []sql.NullBool
	Op     OperatorToSqler
}

func (e ExprMultiNullBool) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiNullBool) ExprValue() ExprMultiValue[sql.NullBool] {
	return ExprMultiValue[sql.NullBool](e)
}

type ExprOr []Where

func (e ExprOr) ToSql() (string, []any, error) {
	if len(e) == 0 {
		return "", []any{}, nil
	}

	b := new(bytes.Buffer)
	vs := make([]any, 0, 10)
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
	Op     OperatorToSqler
}

func (e ExprUint8) ToSql() (string, []any, error) {
	s := strconv.FormatUint(uint64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprUint8) ExprValue() ExprValue[uint8] {
	return ExprValue[uint8](e)
}

type ExprMultiUint8 struct {
	Column string
	Values []uint8
	Op     OperatorToSqler
}

func (e ExprMultiUint8) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiUint8) ExprValue() ExprMultiValue[uint8] {
	return ExprMultiValue[uint8](e)
}

type ExprInt8 struct {
	Column string
	Value  int8
	Op     OperatorToSqler
}

func (e ExprInt8) ToSql() (string, []any, error) {
	s := strconv.FormatInt(int64(e.Value), 10)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{s}, nil
}

func (e ExprInt8) ExprValue() ExprValue[int8] {
	return ExprValue[int8](e)
}

type ExprMultiInt8 struct {
	Column string
	Values []int8
	Op     OperatorToSqler
}

func (e ExprMultiInt8) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiInt8) ExprValue() ExprMultiValue[int8] {
	return ExprMultiValue[int8](e)
}

type ExprBytes struct {
	Column string
	Value  []byte
	Op     OperatorToSqler
}

func (e ExprBytes) ToSql() (string, []any, error) {
	v := sql.RawBytes(e.Value)
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{v}, nil
}

func (e ExprBytes) ExprValue() ExprValue[[]byte] {
	return ExprValue[[]byte](e)
}

type ExprMultiBytes struct {
	Column string
	Values [][]byte
	Op     OperatorToSqler
}

func (e ExprMultiBytes) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(sql.RawBytes(v)))
	}
	return e.Column + " " + ops, vs, nil
}

func (e ExprMultiBytes) ExprValue() ExprMultiValue[[]byte] {
	return ExprMultiValue[[]byte](e)
}

type ExprValue[T any] struct {
	Column string
	Value  T
	Op     OperatorToSqler
}

func (e ExprValue[T]) ToSql() (string, []any, error) {
	v := e.Value
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + " " + ops + " ?", []any{v}, nil
}

func (e ExprValue[T]) ToSqlPg(num int) (string, []any, error) {
	v := e.Value
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	return e.Column + ops + " $" + strconv.Itoa(num), []any{v}, nil
}

type ExprMultiValue[T any] struct {
	Column string
	Values []T
	Op     OperatorToSqler
}

func (e ExprMultiValue[T]) ToSql() (string, []any, error) {
	ops, err := e.Op.ToSql()
	if err != nil {
		return "", nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, vs, nil
}

var ErrNotSupportPg = errors.New("not support pg operator")

func (e ExprMultiValue[T]) ToSqlPg(offset int) (string, int, []any, error) {
	pgop, ok := e.Op.(OperatorToSQLPger)
	if !ok {
		return "", 0, nil, ErrNotSupportPg
	}
	ops, num, err := pgop.ToSqlPg(offset)
	if err != nil {
		return "", 0, nil, err
	}
	vs := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		vs = append(vs, any(v))
	}
	return e.Column + " " + ops, num, vs, nil
}

type ExprNull[T any] struct {
	Column string
	Value  sql.Null[T]
	Op     OperatorToSqler
}

func (e ExprNull[T]) ToSql() (string, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
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

func (e ExprNull[T]) ToSqlPg(offset int) (string, int, []any, error) {
	var ops, placeholder string
	var err error
	vs := []any{}
	if !e.Value.Valid {
		if e.Op == OpNot {
			ops, err = opIsNotNull.ToSql()
		} else {
			ops, err = opIsNull.ToSql()
		}
	} else {
		ops, err = e.Op.ToSql()
		placeholder = " $" + strconv.Itoa(offset+1)
		vs = append(vs, e.Value)
	}
	if err != nil {
		return "", 0, nil, err
	}

	return e.Column + " " + ops + placeholder, offset + 1, vs, nil
}
