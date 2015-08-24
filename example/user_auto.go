package example

import (
	"strings"
	"strconv"

	"github.com/mackee/go-sqlla"
)

type userSQL struct {
	where sqlla.Where
}

func NewUserSQL() userSQL {
	q := userSQL{}
	return q
}


type userSelectSQL struct {
	userSQL
	Columns []string
	order   string
	limit   *uint64
}

func (q userSQL) Select() userSelectSQL {
	return userSelectSQL{
		q,
		[]string{
			"id","name",
		},
		"",
		nil,
	}
}

func (q userSelectSQL) Limit(l uint64) userSelectSQL {
	q.limit = &l
	return q
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
	q.order = " ORDER BY id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
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
	q.order = " ORDER BY name"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM user"
	if wheres != "" {
		query += " WHERE" + wheres
	}
	query += q.order
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}

	return query + ";", vs, nil
}


type userUpdateSQL struct {
	userSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q userSQL) Update() userUpdateSQL {
	return userUpdateSQL{
		userSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q userUpdateSQL) SetID(v uint64) userUpdateSQL {
	q.setMap["id"] = v
	return q
}

func (q userUpdateSQL) WhereID(v uint64, exprs ...sqlla.Operator) userUpdateSQL {
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


func (q userUpdateSQL) SetName(v string) userUpdateSQL {
	q.setMap["name"] = v
	return q
}

func (q userUpdateSQL) WhereName(v string, exprs ...sqlla.Operator) userUpdateSQL {
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


func (q userUpdateSQL) ToSql() (string, []interface{}, error) {
	setColumns, svs, err := q.setMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, wvs, err := q.where.ToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE user SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}


type userInsertSQL struct {
	userSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q userSQL) Insert() userInsertSQL {
	return userInsertSQL{
		userSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q userInsertSQL) ValueID(v uint64) userInsertSQL {
	q.setMap["id"] = v
	return q
}


func (q userInsertSQL) ValueName(v string) userInsertSQL {
	q.setMap["name"] = v
	return q
}


func (q userInsertSQL) ToSql() (string, []interface{}, error) {
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO user " + qs

	return query + ";", vs, nil
}

