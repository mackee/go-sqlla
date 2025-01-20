// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla - DO NOT EDIT.
package example

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/mackee/go-sqlla/v2"
)

type userSQL struct {
	where sqlla.Where
}

func NewUserSQL() userSQL {
	q := userSQL{}
	return q
}

var userAllColumns = []string{
	"`id`", "`name`", "`age`", "`rate`", "`icon_image`", "`created_at`", "`updated_at`",
}

type userSelectSQL struct {
	userSQL
	Columns     []string
	order       sqlla.OrderWithColumn
	limit       *uint64
	offset      *uint64
	tableAlias  string
	joinClauses []string

	additionalWhereClause     string
	additionalWhereClauseArgs []interface{}
	groupByColumns            []string

	isForUpdate bool
}

func (q userSQL) Select() userSelectSQL {
	return userSelectSQL{
		q,
		userAllColumns,
		nil,
		nil,
		nil,
		"",
		nil,
		"",
		nil,

		nil,
		false,
	}
}

func (q userSelectSQL) Or(qs ...userSelectSQL) userSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q userSelectSQL) Limit(l uint64) userSelectSQL {
	q.limit = &l
	return q
}

func (q userSelectSQL) Offset(o uint64) userSelectSQL {
	q.offset = &o
	return q
}

func (q userSelectSQL) ForUpdate() userSelectSQL {
	q.isForUpdate = true
	return q
}

func (q userSelectSQL) TableAlias(alias string) userSelectSQL {
	q.tableAlias = "`" + alias + "`"
	return q
}

func (q userSelectSQL) SetColumns(columns ...string) userSelectSQL {
	q.Columns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(."+"`") {
			q.Columns = append(q.Columns, column)
		} else {
			q.Columns = append(q.Columns, "`"+column+"`")
		}
	}
	return q
}

func (q userSelectSQL) JoinClause(clause string) userSelectSQL {
	q.joinClauses = append(q.joinClauses, clause)
	return q
}

func (q userSelectSQL) AdditionalWhereClause(clause string, args ...interface{}) userSelectSQL {
	q.additionalWhereClause = clause
	q.additionalWhereClauseArgs = args
	return q
}

func (q userSelectSQL) appendColumnPrefix(column string) string {
	if q.tableAlias == "" || strings.ContainsAny(column, "(.") {
		return column
	}
	return q.tableAlias + "." + column
}

func (q userSelectSQL) GroupBy(columns ...string) userSelectSQL {
	q.groupByColumns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(."+"`") {
			q.groupByColumns = append(q.groupByColumns, column)
		} else {
			q.groupByColumns = append(q.groupByColumns, "`"+column+"`")
		}
	}
	return q
}

func (q userSelectSQL) ID(v UserId, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprValue[uint64]{Value: uint64(v), Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) IDIn(vs ...UserId) userSelectSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiValue[uint64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) userSelectSQL {
	v := UserId(pk)
	return q.ID(v, exprs...)
}

func (q userSelectSQL) OrderByID(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`id`"))
	return q
}

func (q userSelectSQL) Name(v string, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`name`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) NameIn(vs ...string) userSelectSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`name`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByName(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`name`"))
	return q
}

func (q userSelectSQL) Age(v sql.NullInt64, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: v.Valid, V: v.Int64}, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`age`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) AgeIn(vs ...sql.NullInt64) userSelectSQL {
	where := sqlla.ExprMultiValue[sql.NullInt64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`age`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByAge(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`age`"))
	return q
}

func (q userSelectSQL) Rate(v float64, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprValue[float64]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`rate`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) RateIn(vs ...float64) userSelectSQL {
	where := sqlla.ExprMultiValue[float64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`rate`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByRate(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`rate`"))
	return q
}

func (q userSelectSQL) IconImage(v []byte, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprValue[[]byte]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`icon_image`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) IconImageIn(vs ...[]byte) userSelectSQL {
	where := sqlla.ExprMultiValue[[]byte]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`icon_image`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByIconImage(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`icon_image`"))
	return q
}

func (q userSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) CreatedAtIn(vs ...time.Time) userSelectSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByCreatedAt(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`created_at`"))
	return q
}

func (q userSelectSQL) UpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userSelectSQL {
	where := sqlla.ExprNull[time.Time]{Value: sql.Null[time.Time]{Valid: v.Valid, V: v.Time}, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) UpdatedAtIn(vs ...mysql.NullTime) userSelectSQL {
	where := sqlla.ExprMultiValue[mysql.NullTime]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q userSelectSQL) OrderByUpdatedAt(order sqlla.Order) userSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("`updated_at`"))
	return q
}

func (q userSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	tableName := "`user`"
	if q.tableAlias != "" {
		tableName = tableName + " AS " + q.tableAlias
		pcs := make([]string, 0, len(q.Columns))
		for _, column := range q.Columns {
			pcs = append(pcs, q.appendColumnPrefix(column))
		}
		columns = strings.Join(pcs, ", ")
	}
	query := "SELECT " + columns + " FROM " + tableName
	if len(q.joinClauses) > 0 {
		jc := strings.Join(q.joinClauses, " ")
		query += " " + jc
	}
	if wheres != "" {
		query += " WHERE" + wheres
	}
	if q.additionalWhereClause != "" {
		query += " " + q.additionalWhereClause
		if len(q.additionalWhereClauseArgs) > 0 {
			vs = append(vs, q.additionalWhereClauseArgs...)
		}
	}
	if len(q.groupByColumns) > 0 {
		query += " GROUP BY "
		gbcs := make([]string, 0, len(q.groupByColumns))
		for _, column := range q.groupByColumns {
			gbcs = append(gbcs, q.appendColumnPrefix(column))
		}
		query += strings.Join(gbcs, ", ")
	}
	if q.order != nil {
		query += " ORDER BY " + q.order.OrderExpr()
		vs = append(vs, q.order.Values()...)
	}
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}
	if q.offset != nil {
		query += " OFFSET " + strconv.FormatUint(*q.offset, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (s User) Select() userSelectSQL {
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

func (q userSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (User, error) {
	q.Columns = userAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
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

func (q userSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]User, error) {
	rs := make([]User, 0, 10)
	q.Columns = userAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
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
		&row.IconImage,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	return row, err
}

// IterContext returns iter.Seq2[User, error] and closer.
//
// The returned Iter.Seq2 assembles and executes a query in the first iteration.
// Therefore, the first iteration may return an error in assembling or executing the query.
// Subsequent iterations read rows. Again, the read may return an error.
//
// closer is a function that closes the row reader object. Execution of this function is idempotent.
// Be sure to call it when you are done using iter.Seq2.
func (q userSelectSQL) IterContext(ctx context.Context, db sqlla.DB) (func(func(User, error) bool), func() error) {
	var rowClose func() error
	closer := func() error {
		if rowClose != nil {
			err := rowClose()
			rowClose = nil
			return err
		}
		return nil
	}

	q.Columns = userAllColumns
	query, args, err := q.ToSql()
	return func(yield func(User, error) bool) {
		if err != nil {
			var r User
			yield(r, err)
			return
		}
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			var r User
			yield(r, err)
			return
		}
		rowClose = rows.Close
		for rows.Next() {
			r, err := q.Scan(rows)
			if !yield(r, err) {
				break
			}
		}
	}, closer
}

type userUpdateSQL struct {
	userSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q userSQL) Update() userUpdateSQL {
	return userUpdateSQL{
		userSQL: q,
		setMap:  sqlla.SetMap{},
	}
}

func (q userUpdateSQL) SetID(v UserId) userUpdateSQL {
	q.setMap["`id`"] = uint64(v)
	return q
}

func (q userUpdateSQL) WhereID(v UserId, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprValue[uint64]{Value: uint64(v), Op: sqlla.Operators(exprs), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereIDIn(vs ...UserId) userUpdateSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiValue[uint64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetName(v string) userUpdateSQL {
	q.setMap["`name`"] = v
	return q
}

func (q userUpdateSQL) WhereName(v string, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereNameIn(vs ...string) userUpdateSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetAge(v sql.NullInt64) userUpdateSQL {
	q.setMap["`age`"] = sql.Null[int64]{Valid: v.Valid, V: v.Int64}
	return q
}

func (q userUpdateSQL) WhereAge(v sql.NullInt64, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: v.Valid, V: v.Int64}, Op: sqlla.Operators(exprs), Column: "`age`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereAgeIn(vs ...sql.NullInt64) userUpdateSQL {
	where := sqlla.ExprMultiValue[sql.NullInt64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`age`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetRate(v float64) userUpdateSQL {
	q.setMap["`rate`"] = v
	return q
}

func (q userUpdateSQL) WhereRate(v float64, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprValue[float64]{Value: v, Op: sqlla.Operators(exprs), Column: "`rate`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereRateIn(vs ...float64) userUpdateSQL {
	where := sqlla.ExprMultiValue[float64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`rate`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetIconImage(v []byte) userUpdateSQL {
	q.setMap["`icon_image`"] = v
	return q
}

func (q userUpdateSQL) WhereIconImage(v []byte, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprValue[[]byte]{Value: v, Op: sqlla.Operators(exprs), Column: "`icon_image`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereIconImageIn(vs ...[]byte) userUpdateSQL {
	where := sqlla.ExprMultiValue[[]byte]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`icon_image`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetCreatedAt(v time.Time) userUpdateSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q userUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereCreatedAtIn(vs ...time.Time) userUpdateSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) SetUpdatedAt(v mysql.NullTime) userUpdateSQL {
	q.setMap["`updated_at`"] = sql.Null[time.Time]{Valid: v.Valid, V: v.Time}
	return q
}

func (q userUpdateSQL) WhereUpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userUpdateSQL {
	where := sqlla.ExprNull[time.Time]{Value: sql.Null[time.Time]{Valid: v.Valid, V: v.Time}, Op: sqlla.Operators(exprs), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userUpdateSQL) WhereUpdatedAtIn(vs ...mysql.NullTime) userUpdateSQL {
	where := sqlla.ExprMultiValue[mysql.NullTime]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
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

	query := "UPDATE " + "`user`" + " SET" + setColumns
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

func (q userUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]User, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userSQL

	return qq.Select().AllContext(ctx, db)
}

type userDefaultUpdateHooker interface {
	DefaultUpdateHook(userUpdateSQL) (userUpdateSQL, error)
}

type userInsertSQL struct {
	userSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q userSQL) Insert() userInsertSQL {
	return userInsertSQL{
		userSQL: q,
		setMap:  sqlla.SetMap{},
	}
}

func (q userInsertSQL) ValueID(v UserId) userInsertSQL {
	q.setMap["`id`"] = uint64(v)
	return q
}

func (q userInsertSQL) ValueName(v string) userInsertSQL {
	q.setMap["`name`"] = v
	return q
}

func (q userInsertSQL) ValueAge(v sql.NullInt64) userInsertSQL {
	q.setMap["`age`"] = sql.Null[int64]{Valid: v.Valid, V: v.Int64}
	return q
}

func (q userInsertSQL) ValueRate(v float64) userInsertSQL {
	q.setMap["`rate`"] = v
	return q
}

func (q userInsertSQL) ValueIconImage(v []byte) userInsertSQL {
	q.setMap["`icon_image`"] = v
	return q
}

func (q userInsertSQL) ValueCreatedAt(v time.Time) userInsertSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q userInsertSQL) ValueUpdatedAt(v mysql.NullTime) userInsertSQL {
	q.setMap["`updated_at`"] = sql.Null[time.Time]{Valid: v.Valid, V: v.Time}
	return q
}

func (q userInsertSQL) ToSql() (string, []any, error) {
	query, vs, err := q.userInsertSQLToSql()
	if err != nil {
		return "", []any{}, err
	}
	return query + ";", vs, nil
}

func (q userInsertSQL) userInsertSQLToSql() (string, []any, error) {
	var err error
	var s interface{} = User{}
	if t, ok := s.(userDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []any{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []any{}, err
	}

	query := "INSERT INTO " + "`user`" + " " + qs
	return query, vs, nil
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

func (q userInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (User, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return NewUserSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q userInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userDefaultInsertHooker interface {
	DefaultInsertHook(userInsertSQL) (userInsertSQL, error)
}

type userInsertSQLToSqler interface {
	userInsertSQLToSql() (string, []any, error)
}

type userBulkInsertSQL struct {
	insertSQLs []userInsertSQL
}

func (q userSQL) BulkInsert() *userBulkInsertSQL {
	return &userBulkInsertSQL{
		insertSQLs: []userInsertSQL{},
	}
}

func (q *userBulkInsertSQL) Append(iqs ...userInsertSQL) {
	q.insertSQLs = append(q.insertSQLs, iqs...)
}

func (q *userBulkInsertSQL) userInsertSQLToSql() (string, []any, error) {
	if len(q.insertSQLs) == 0 {
		return "", []any{}, fmt.Errorf("sqlla: This userBulkInsertSQL's InsertSQL was empty")
	}
	iqs := make([]userInsertSQL, len(q.insertSQLs))
	copy(iqs, q.insertSQLs)

	var s interface{} = User{}
	if t, ok := s.(userDefaultInsertHooker); ok {
		for i, iq := range iqs {
			var err error
			iq, err = t.DefaultInsertHook(iq)
			if err != nil {
				return "", []any{}, err
			}
			iqs[i] = iq
		}
	}

	sms := make(sqlla.SetMaps, 0, len(q.insertSQLs))
	for _, iq := range q.insertSQLs {
		sms = append(sms, iq.setMap)
	}

	query, vs, err := sms.ToInsertSql()
	if err != nil {
		return "", []any{}, err
	}
	return "INSERT INTO " + "`user`" + " " + query, vs, nil
}

func (q *userBulkInsertSQL) ToSql() (string, []any, error) {
	query, vs, err := q.userInsertSQLToSql()
	if err != nil {
		return "", []any{}, err
	}
	return query + ";", vs, nil
}
func (q *userBulkInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userInsertOnDuplicateKeyUpdateSQL struct {
	insertSQL               userInsertSQLToSqler
	onDuplicateKeyUpdateMap sqlla.SetMap
}

func (q userInsertSQL) OnDuplicateKeyUpdate() userInsertOnDuplicateKeyUpdateSQL {
	return userInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateID(v UserId) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = uint64(v)
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateID() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = sqlla.SetMapRawValue("VALUES(" + "`id`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateName(v string) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateName(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateName() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = sqlla.SetMapRawValue("VALUES(" + "`name`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateAge(v sql.NullInt64) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`age`"] = sql.Null[int64]{Valid: v.Valid, V: v.Int64}
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateAge(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`age`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateAge() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`age`"] = sqlla.SetMapRawValue("VALUES(" + "`age`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateRate(v float64) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`rate`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateRate(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`rate`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateRate() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`rate`"] = sqlla.SetMapRawValue("VALUES(" + "`rate`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateIconImage(v []byte) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`icon_image`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateIconImage(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`icon_image`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateIconImage() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`icon_image`"] = sqlla.SetMapRawValue("VALUES(" + "`icon_image`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateCreatedAt(v time.Time) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateCreatedAt() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = sqlla.SetMapRawValue("VALUES(" + "`created_at`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateUpdatedAt(v mysql.NullTime) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = sql.Null[time.Time]{Valid: v.Valid, V: v.Time}
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateUpdatedAt(v sqlla.SetMapRawValue) userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = v
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) SameOnUpdateUpdatedAt() userInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = sqlla.SetMapRawValue("VALUES(" + "`updated_at`" + ")")
	return q
}

func (q userInsertOnDuplicateKeyUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = User{}
	if t, ok := s.(userDefaultInsertOnDuplicateKeyUpdateHooker); ok {
		q, err = t.DefaultInsertOnDuplicateKeyUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}

	query, vs, err := q.insertSQL.userInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	os, ovs, err := q.onDuplicateKeyUpdateMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	query += " ON DUPLICATE KEY UPDATE" + os
	vs = append(vs, ovs...)

	return query + ";", vs, nil
}

func (q userInsertOnDuplicateKeyUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) (User, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return User{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}
	return NewUserSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q userInsertOnDuplicateKeyUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type userDefaultInsertOnDuplicateKeyUpdateHooker interface {
	DefaultInsertOnDuplicateKeyUpdateHook(userInsertOnDuplicateKeyUpdateSQL) (userInsertOnDuplicateKeyUpdateSQL, error)
}

func (q *userBulkInsertSQL) OnDuplicateKeyUpdate() userInsertOnDuplicateKeyUpdateSQL {
	return userInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

type userDeleteSQL struct {
	userSQL
}

func (q userSQL) Delete() userDeleteSQL {
	return userDeleteSQL{
		q,
	}
}

func (q userDeleteSQL) ID(v UserId, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprValue[uint64]{Value: uint64(v), Op: sqlla.Operators(exprs), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) IDIn(vs ...UserId) userDeleteSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiValue[uint64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) Name(v string, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) NameIn(vs ...string) userDeleteSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) Age(v sql.NullInt64, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: v.Valid, V: v.Int64}, Op: sqlla.Operators(exprs), Column: "`age`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) AgeIn(vs ...sql.NullInt64) userDeleteSQL {
	where := sqlla.ExprMultiValue[sql.NullInt64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`age`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) Rate(v float64, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprValue[float64]{Value: v, Op: sqlla.Operators(exprs), Column: "`rate`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) RateIn(vs ...float64) userDeleteSQL {
	where := sqlla.ExprMultiValue[float64]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`rate`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) IconImage(v []byte, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprValue[[]byte]{Value: v, Op: sqlla.Operators(exprs), Column: "`icon_image`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) IconImageIn(vs ...[]byte) userDeleteSQL {
	where := sqlla.ExprMultiValue[[]byte]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`icon_image`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) CreatedAtIn(vs ...time.Time) userDeleteSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) UpdatedAt(v mysql.NullTime, exprs ...sqlla.Operator) userDeleteSQL {
	where := sqlla.ExprNull[time.Time]{Value: sql.Null[time.Time]{Valid: v.Valid, V: v.Time}, Op: sqlla.Operators(exprs), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) UpdatedAtIn(vs ...mysql.NullTime) userDeleteSQL {
	where := sqlla.ExprMultiValue[mysql.NullTime]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q userDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM " + "`user`"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func (q userDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (q userDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s User) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s User) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
