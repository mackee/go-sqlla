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

var allColumns = []string{
	"id","name",
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
		allColumns,
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

func (q userSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) userSelectSQL {
	v := uint64(pk)
	return q.ID(v, exprs...)
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

func (q userSelectSQL) Single(db sqlla.DB) (User, error) {
	q.Columns = allColumns
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q userSelectSQL) All(db sqlla.DB) ([]User, error) {
	rs := make([]User, 0, 10)
	q.Columns = allColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q userSelectSQL) Scan(s sqlla.Scanner) (User, error) {
	var row User
	err := s.Scan(
		&row.Id,
		&row.Name,
		
	)
	return row, err
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

func (q userInsertSQL) Exec(db sqlla.DB) (User, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return NewUserSQL().Select().PkColumn(id).Single(db)
}


type userDeleteSQL struct {
	userSQL
}

func (q userSQL) Delete() userDeleteSQL {
	return userDeleteSQL{
		q,
	}
}


func (q userDeleteSQL) ID(v uint64, exprs ...sqlla.Operator) userDeleteSQL {
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


func (q userDeleteSQL) Name(v string, exprs ...sqlla.Operator) userDeleteSQL {
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


func (q userDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM user"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

