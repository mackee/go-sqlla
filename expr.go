package sqlla

import (
	"bytes"
	"database/sql"
	"errors"
	"strconv"
)

type Expr interface {
	ToSql() (string, []any, error)
	ToSqlPg(offset int) (string, int, []any, error)
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

func (e ExprOr) ToSqlPg(offset int) (string, int, []any, error) {
	if len(e) == 0 {
		return "", offset, []any{}, nil
	}

	b := new(bytes.Buffer)
	vs := make([]any, 0, 10)
	b.WriteString("(")
	for i, w := range e {
		if i > 0 {
			b.WriteString(" OR ")
		}
		b.WriteString("(")
		q, n, wvs, err := w.ToSqlPg(offset)
		if err != nil {
			return "", 0, nil, err
		}
		b.WriteString(q)
		b.WriteString(" )")
		vs = append(vs, wvs...)
		offset = n
	}

	b.WriteString(")")

	return b.String(), offset, vs, nil
}

type ExprValue[T any] struct {
	Column string
	Value  T
	Op     Operators
}

func (e ExprValue[T]) ToSql() (string, []any, error) {
	o := e.Op
	var o1 Operator = OpEqual
	if len(o) > 0 {
		o1, o = o[0], o[1:]
	}
	o1op, err := o1.ToSql()
	if err != nil {
		return "", nil, err
	}
	expr := e.Column + " " + o1op + " ?"
	for _, op := range o {
		on, err := op.ToSql()
		if err != nil {
			return "", nil, err
		}
		expr = "( " + expr + " ) " + on + " ?"
	}
	return expr, append([]any{e.Value}, o.Values()...), nil
}

func (e ExprValue[T]) ToSqlPg(offset int) (string, int, []any, error) {
	o := e.Op
	var o1 Operator = OpEqual
	if len(o) > 0 {
		o1, o = o[0], o[1:]
	}
	o1op, err := o1.ToSql()
	if err != nil {
		return "", 0, nil, err
	}
	offset++
	expr := e.Column + " " + o1op + " $" + strconv.Itoa(offset)
	for _, op := range o {
		on, err := op.ToSql()
		if err != nil {
			return "", 0, nil, err
		}
		offset++
		expr = "( " + expr + " ) " + on + " $" + strconv.Itoa(offset)
	}
	return expr, offset, append([]any{e.Value}, o.Values()...), nil
}

type ExprMultiValue[T any] struct {
	Column string
	Values []T
	Op     OperatorMulti
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
	ops, num, err := e.Op.ToSqlPg(offset)
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
	Op     Operators
}

func (e ExprNull[T]) ToSql() (string, []any, error) {
	var expr string
	o := e.Op
	vs := make([]any, 0, len(e.Op))
	if !e.Value.Valid {
		var o1 Operator = opIsNull
		if len(o) > 0 {
			if o[0] == OpNot {
				o1 = opIsNotNull
			}
			o = o[1:]
		}
		ops, err := o1.ToSql()
		if err != nil {
			return "", nil, err
		}
		expr = e.Column + " " + ops
	} else {
		var o1 Operator = OpEqual
		if len(e.Op) > 0 {
			o1, o = o[0], o[1:]
		}
		ops, err := o1.ToSql()
		if err != nil {
			return "", nil, err
		}
		expr = e.Column + " " + ops + " ?"
		vs = append(vs, e.Value.V)
	}
	for _, op := range o {
		on, err := op.ToSql()
		if err != nil {
			return "", nil, err
		}
		expr = "( " + expr + " ) " + on + " ?"
	}
	return expr, append(vs, o.Values()...), nil
}

func (e ExprNull[T]) ToSqlPg(offset int) (string, int, []any, error) {
	var expr string
	o := e.Op
	vs := make([]any, 0, len(e.Op))
	if !e.Value.Valid {
		var o1 Operator = opIsNull
		if len(o) > 0 {
			if o[0] == OpNot {
				o1 = opIsNotNull
			}
			o = o[1:]
		}
		ops, err := o1.ToSql()
		if err != nil {
			return "", 0, nil, err
		}
		expr = e.Column + " " + ops
	} else {
		var o1 Operator = OpEqual
		if len(e.Op) > 0 {
			o1, o = e.Op[0], e.Op[1:]
		}
		ops, err := o1.ToSql()
		if err != nil {
			return "", 0, nil, err
		}
		offset++
		expr = e.Column + " " + ops + " $" + strconv.Itoa(offset)
		vs = append(vs, e.Value.V)
	}
	for _, op := range o {
		on, err := op.ToSql()
		if err != nil {
			return "", 0, nil, err
		}
		offset++
		expr = "( " + expr + " ) " + on + " $" + strconv.Itoa(offset)
	}
	return expr, offset, append(vs, o.Values()...), nil
}
