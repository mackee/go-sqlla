package example

import (
	"strings"

	"github.com/mackee/go-sqlla"
)

type userSQL struct {
	where []sqlla.Expr
}

type userSelectSQL struct {
	userSQL
	Columns []string
	Order   string
}

func NewUserSQL() userSQL {
	q := userSQL{}
	return q
}

func (q userSQL) Select() userSelectSQL {
	return userSelectSQL{
		q,
		[]string{
			"id","name",
		},
		"",
	}
}


func (q userSelectSQL) ID(v uint64, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByID(order sqlla.Order) userSelectSQL {
	q.Order = " ORDER BY id"
	if order == sqlla.Asc {
		q.Order += " ASC"
	} else {
		q.Order += " DESC"
	}

	return q
}

func (q userSelectSQL) Name(v string, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "name"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByName(order sqlla.Order) userSelectSQL {
	q.Order = " ORDER BY name"
	if order == sqlla.Asc {
		q.Order += " ASC"
	} else {
		q.Order += " DESC"
	}

	return q
}

func (q userSelectSQL) ToSelectSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	vs := []interface{}{}
	var wheres string
	for i, w := range q.where {
		s, v, err := w.ToSql()
		if err != nil {
			return "", nil, err
		}
		vs = append(vs, v...)

		if i == 0 {
			wheres += s
			continue
		}
		wheres += " AND " + s
	}

	return "SELECT " + columns + " FROM user WHERE " + wheres + q.Order + ";", vs, nil
}
