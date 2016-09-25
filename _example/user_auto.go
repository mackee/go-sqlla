package example

import (
	"strings"
	"strconv"

	"database/sql"
	"time"
	"github.com/go-sql-driver/mysql"
	
	"github.com/mackee/go-sqlla"
)

type userSQL struct {
	where sqlla.Where
}

func NewUserSQL() userSQL {
	q := userSQL{}
	return q
}

var userAllColumns = []string{
	"id","name","age","rate","created_at","updated_at",
}

type userSelectSQL struct {
	userSQL
	Columns     []string
	order       string
	limit       *uint64
	isForUpdate bool
}

func (q userSQL) Select() userSelectSQL {
	return userSelectSQL{
		q,
		userAllColumns,
		"",
		nil,
		false,
	}
}

func (q userSelectSQL) Limit(l uint64) userSelectSQL {
	q.limit = &l
	return q
}

func (q userSelectSQL) ForUpdate() userSelectSQL {
	q.isForUpdate = true
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

func (q userSelectSQL) IDIn(vs ...uint64) userSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
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

func (q userSelectSQL) NameIn(vs ...string) userSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "name"}
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

func (q userSelectSQL) Age(v sql.NullInt64, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "age"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) AgeIn(vs ...sql.NullInt64) userSelectSQL {
	where := sqlla.ExprMultiNullInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "age"}
	q.where = append(q.where, where)
	return q
}



func (q userSelectSQL) OrderByAge(order sqlla.Order) userSelectSQL {
	q.order = " ORDER BY age"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userSelectSQL) Rate(v float64, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "rate"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) RateIn(vs ...float64) userSelectSQL {
	where := sqlla.ExprMultiFloat64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "rate"}
	q.where = append(q.where, where)
	return q
}



func (q userSelectSQL) OrderByRate(order sqlla.Order) userSelectSQL {
	q.order = " ORDER BY rate"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) CreatedAtIn(vs ...time.Time) userSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created_at"}
	q.where = append(q.where, where)
	return q
}



func (q userSelectSQL) OrderByCreatedAt(order sqlla.Order) userSelectSQL {
	q.order = " ORDER BY created_at"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userSelectSQL) UpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullTime{Value: v, Op: op, Column: "updated_at"}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) UpdatedAtIn(vs ...mysql.NullTime) userSelectSQL {
	where := sqlla.ExprMultiNullTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "updated_at"}
	q.where = append(q.where, where)
	return q
}



func (q userSelectSQL) OrderByUpdatedAt(order sqlla.Order) userSelectSQL {
	q.order = " ORDER BY updated_at"
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

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (s User) Select() (userSelectSQL) {
	return NewUserSQL().Select().ID(s.Id)
}
func (q userSelectSQL) Single(db sqlla.DB) (User, error) {
	q.Columns = userAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q userSelectSQL) All(db sqlla.DB) ([]User, error) {
	rs := make([]User, 0, 10)
	q.Columns = userAllColumns
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
		&row.Age,
		&row.Rate,
		&row.CreatedAt,
		&row.UpdatedAt,
		
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


func (q userUpdateSQL) SetAge(v sql.NullInt64) userUpdateSQL {
	q.setMap["age"] = v
	return q
}

func (q userUpdateSQL) WhereAge(v sql.NullInt64, exprs ...sqlla.Operator) userUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "age"}
	q.where = append(q.where, where)
	return q
}


func (q userUpdateSQL) SetRate(v float64) userUpdateSQL {
	q.setMap["rate"] = v
	return q
}

func (q userUpdateSQL) WhereRate(v float64, exprs ...sqlla.Operator) userUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "rate"}
	q.where = append(q.where, where)
	return q
}


func (q userUpdateSQL) SetCreatedAt(v time.Time) userUpdateSQL {
	q.setMap["created_at"] = v
	return q
}

func (q userUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) userUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}


func (q userUpdateSQL) SetUpdatedAt(v mysql.NullTime) userUpdateSQL {
	q.setMap["updated_at"] = v
	return q
}

func (q userUpdateSQL) WhereUpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullTime{Value: v, Op: op, Column: "updated_at"}
	q.where = append(q.where, where)
	return q
}


func (q userUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = User{}
	if t, ok := s.(userDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
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
func (s User) Update() userUpdateSQL {
	return NewUserSQL().Update().WhereID(s.Id)
}

func (q userUpdateSQL) Exec(db sqlla.DB) ([]User, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userSQL

	return qq.Select().All(db)
}

type userDefaultUpdateHooker interface {
	DefaultUpdateHook(userUpdateSQL) (userUpdateSQL, error)
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


func (q userInsertSQL) ValueAge(v sql.NullInt64) userInsertSQL {
	q.setMap["age"] = v
	return q
}


func (q userInsertSQL) ValueRate(v float64) userInsertSQL {
	q.setMap["rate"] = v
	return q
}


func (q userInsertSQL) ValueCreatedAt(v time.Time) userInsertSQL {
	q.setMap["created_at"] = v
	return q
}


func (q userInsertSQL) ValueUpdatedAt(v mysql.NullTime) userInsertSQL {
	q.setMap["updated_at"] = v
	return q
}


func (q userInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = User{}
	if t, ok := s.(userDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
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

type userDefaultInsertHooker interface {
	DefaultInsertHook(userInsertSQL) (userInsertSQL, error)
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


func (q userDeleteSQL) IDIn(vs ...uint64) userDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
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


func (q userDeleteSQL) NameIn(vs ...string) userDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "name"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) Age(v sql.NullInt64, exprs ...sqlla.Operator) userDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullInt64{Value: v, Op: op, Column: "age"}
	q.where = append(q.where, where)
	return q
}


func (q userDeleteSQL) AgeIn(vs ...sql.NullInt64) userDeleteSQL {
	where := sqlla.ExprMultiNullInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "age"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) Rate(v float64, exprs ...sqlla.Operator) userDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprFloat64{Value: v, Op: op, Column: "rate"}
	q.where = append(q.where, where)
	return q
}


func (q userDeleteSQL) RateIn(vs ...float64) userDeleteSQL {
	where := sqlla.ExprMultiFloat64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "rate"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprTime{Value: v, Op: op, Column: "created_at"}
	q.where = append(q.where, where)
	return q
}


func (q userDeleteSQL) CreatedAtIn(vs ...time.Time) userDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "created_at"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) UpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullTime{Value: v, Op: op, Column: "updated_at"}
	q.where = append(q.where, where)
	return q
}


func (q userDeleteSQL) UpdatedAtIn(vs ...mysql.NullTime) userDeleteSQL {
	where := sqlla.ExprMultiNullTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "updated_at"}
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

func ( q userDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}
func (s User) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

