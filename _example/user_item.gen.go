package example

import (
	"strings"
	"strconv"
	"context"

	"database/sql"
	
	"github.com/mackee/go-sqlla"
)

type userItemSQL struct {
	where sqlla.Where
}

func NewUserItemSQL() userItemSQL {
	q := userItemSQL{}
	return q
}

var userItemAllColumns = []string{
	"id","user_id","item_id","is_used","has_extension",
}

type userItemSelectSQL struct {
	userItemSQL
	Columns     []string
	order       string
	limit       *uint64
	offset      *uint64
	isForUpdate bool
}

func (q userItemSQL) Select() userItemSelectSQL {
	return userItemSelectSQL{
		q,
		userItemAllColumns,
		"",
		nil,
		nil,
		false,
	}
}

func (q userItemSelectSQL) Or(qs ...userItemSelectSQL) userItemSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q userItemSelectSQL) Limit(l uint64) userItemSelectSQL {
	q.limit = &l
	return q
}

func (q userItemSelectSQL) Offset(o uint64) userItemSelectSQL {
	q.offset = &o
	return q
}

func (q userItemSelectSQL) ForUpdate() userItemSelectSQL {
	q.isForUpdate = true
	return q
}


func (q userItemSelectSQL) ID(v uint64, exprs ...sqlla.Operator) userItemSelectSQL {
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

func (q userItemSelectSQL) IDIn(vs ...uint64) userItemSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) userItemSelectSQL {
	v := uint64(pk)
	return q.ID(v, exprs...)
}

func (q userItemSelectSQL) OrderByID(order sqlla.Order) userItemSelectSQL {
	q.order = " ORDER BY id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userItemSelectSQL) UserID(v uint64, exprs ...sqlla.Operator) userItemSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "user_id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemSelectSQL) UserIDIn(vs ...uint64) userItemSelectSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "user_id"}
	q.where = append(q.where, where)
	return q
}



func (q userItemSelectSQL) OrderByUserID(order sqlla.Order) userItemSelectSQL {
	q.order = " ORDER BY user_id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userItemSelectSQL) ItemID(v string, exprs ...sqlla.Operator) userItemSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "item_id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemSelectSQL) ItemIDIn(vs ...string) userItemSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "item_id"}
	q.where = append(q.where, where)
	return q
}



func (q userItemSelectSQL) OrderByItemID(order sqlla.Order) userItemSelectSQL {
	q.order = " ORDER BY item_id"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userItemSelectSQL) IsUsed(v bool, exprs ...sqlla.Operator) userItemSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprBool{Value: v, Op: op, Column: "is_used"}
	q.where = append(q.where, where)
	return q
}

func (q userItemSelectSQL) IsUsedIn(vs ...bool) userItemSelectSQL {
	where := sqlla.ExprMultiBool{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "is_used"}
	q.where = append(q.where, where)
	return q
}



func (q userItemSelectSQL) OrderByIsUsed(order sqlla.Order) userItemSelectSQL {
	q.order = " ORDER BY is_used"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userItemSelectSQL) HasExtension(v sql.NullBool, exprs ...sqlla.Operator) userItemSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullBool{Value: v, Op: op, Column: "has_extension"}
	q.where = append(q.where, where)
	return q
}

func (q userItemSelectSQL) HasExtensionIn(vs ...sql.NullBool) userItemSelectSQL {
	where := sqlla.ExprMultiNullBool{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "has_extension"}
	q.where = append(q.where, where)
	return q
}



func (q userItemSelectSQL) OrderByHasExtension(order sqlla.Order) userItemSelectSQL {
	q.order = " ORDER BY has_extension"
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q userItemSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "SELECT " + columns + " FROM user_item"
	if wheres != "" {
		query += " WHERE" + wheres
	}
	query += q.order
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

func (s UserItem) Select() (userItemSelectSQL) {
	return NewUserItemSQL().Select().ID(s.Id)
}
func (q userItemSelectSQL) Single(db sqlla.DB) (UserItem, error) {
	q.Columns = userItemAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return UserItem{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q userItemSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (UserItem, error) {
	q.Columns = userItemAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return UserItem{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
	return q.Scan(row)
}

func (q userItemSelectSQL) All(db sqlla.DB) ([]UserItem, error) {
	rs := make([]UserItem, 0, 10)
	q.Columns = userItemAllColumns
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

func (q userItemSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]UserItem, error) {
	rs := make([]UserItem, 0, 10)
	q.Columns = userItemAllColumns
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

func (q userItemSelectSQL) Scan(s sqlla.Scanner) (UserItem, error) {
	var row UserItem
	err := s.Scan(
		&row.Id,
		&row.UserId,
		&row.ItemId,
		&row.IsUsed,
		&row.HasExtension,
		
	)
	return row, err
}

type userItemUpdateSQL struct {
	userItemSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q userItemSQL) Update() userItemUpdateSQL {
	return userItemUpdateSQL{
		userItemSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q userItemUpdateSQL) SetID(v uint64) userItemUpdateSQL {
	q.setMap["id"] = v
	return q
}

func (q userItemUpdateSQL) WhereID(v uint64, exprs ...sqlla.Operator) userItemUpdateSQL {
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


func (q userItemUpdateSQL) SetUserID(v uint64) userItemUpdateSQL {
	q.setMap["user_id"] = v
	return q
}

func (q userItemUpdateSQL) WhereUserID(v uint64, exprs ...sqlla.Operator) userItemUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "user_id"}
	q.where = append(q.where, where)
	return q
}


func (q userItemUpdateSQL) SetItemID(v string) userItemUpdateSQL {
	q.setMap["item_id"] = v
	return q
}

func (q userItemUpdateSQL) WhereItemID(v string, exprs ...sqlla.Operator) userItemUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "item_id"}
	q.where = append(q.where, where)
	return q
}


func (q userItemUpdateSQL) SetIsUsed(v bool) userItemUpdateSQL {
	q.setMap["is_used"] = v
	return q
}

func (q userItemUpdateSQL) WhereIsUsed(v bool, exprs ...sqlla.Operator) userItemUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprBool{Value: v, Op: op, Column: "is_used"}
	q.where = append(q.where, where)
	return q
}


func (q userItemUpdateSQL) SetHasExtension(v sql.NullBool) userItemUpdateSQL {
	q.setMap["has_extension"] = v
	return q
}

func (q userItemUpdateSQL) WhereHasExtension(v sql.NullBool, exprs ...sqlla.Operator) userItemUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullBool{Value: v, Op: op, Column: "has_extension"}
	q.where = append(q.where, where)
	return q
}


func (q userItemUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = UserItem{}
	if t, ok := s.(userItemDefaultUpdateHooker); ok {
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

	query := "UPDATE user_item SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s UserItem) Update() userItemUpdateSQL {
	return NewUserItemSQL().Update().WhereID(s.Id)
}

func (q userItemUpdateSQL) Exec(db sqlla.DB) ([]UserItem, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userItemSQL

	return qq.Select().All(db)
}

func (q userItemUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]UserItem, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.userItemSQL

	return qq.Select().AllContext(ctx, db)
}

type userItemDefaultUpdateHooker interface {
	DefaultUpdateHook(userItemUpdateSQL) (userItemUpdateSQL, error)
}

type userItemInsertSQL struct {
	userItemSQL
	setMap	sqlla.SetMap
	Columns []string
}

func (q userItemSQL) Insert() userItemInsertSQL {
	return userItemInsertSQL{
		userItemSQL: q,
		setMap: sqlla.SetMap{},
	}
}


func (q userItemInsertSQL) ValueID(v uint64) userItemInsertSQL {
	q.setMap["id"] = v
	return q
}


func (q userItemInsertSQL) ValueUserID(v uint64) userItemInsertSQL {
	q.setMap["user_id"] = v
	return q
}


func (q userItemInsertSQL) ValueItemID(v string) userItemInsertSQL {
	q.setMap["item_id"] = v
	return q
}


func (q userItemInsertSQL) ValueIsUsed(v bool) userItemInsertSQL {
	q.setMap["is_used"] = v
	return q
}


func (q userItemInsertSQL) ValueHasExtension(v sql.NullBool) userItemInsertSQL {
	q.setMap["has_extension"] = v
	return q
}


func (q userItemInsertSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = UserItem{}
	if t, ok := s.(userItemDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO user_item " + qs

	return query + ";", vs, nil
}

func (q userItemInsertSQL) Exec(db sqlla.DB) (UserItem, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return UserItem{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return UserItem{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserItem{}, err
	}
	return NewUserItemSQL().Select().PkColumn(id).Single(db)
}

func (q userItemInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (UserItem, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return UserItem{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return UserItem{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return UserItem{}, err
	}
	return NewUserItemSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

type userItemDefaultInsertHooker interface {
	DefaultInsertHook(userItemInsertSQL) (userItemInsertSQL, error)
}

type userItemDeleteSQL struct {
	userItemSQL
}

func (q userItemSQL) Delete() userItemDeleteSQL {
	return userItemDeleteSQL{
		q,
	}
}


func (q userItemDeleteSQL) ID(v uint64, exprs ...sqlla.Operator) userItemDeleteSQL {
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


func (q userItemDeleteSQL) IDIn(vs ...uint64) userItemDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemDeleteSQL) UserID(v uint64, exprs ...sqlla.Operator) userItemDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprUint64{Value: v, Op: op, Column: "user_id"}
	q.where = append(q.where, where)
	return q
}


func (q userItemDeleteSQL) UserIDIn(vs ...uint64) userItemDeleteSQL {
	where := sqlla.ExprMultiUint64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "user_id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemDeleteSQL) ItemID(v string, exprs ...sqlla.Operator) userItemDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprString{Value: v, Op: op, Column: "item_id"}
	q.where = append(q.where, where)
	return q
}


func (q userItemDeleteSQL) ItemIDIn(vs ...string) userItemDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "item_id"}
	q.where = append(q.where, where)
	return q
}

func (q userItemDeleteSQL) IsUsed(v bool, exprs ...sqlla.Operator) userItemDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprBool{Value: v, Op: op, Column: "is_used"}
	q.where = append(q.where, where)
	return q
}


func (q userItemDeleteSQL) IsUsedIn(vs ...bool) userItemDeleteSQL {
	where := sqlla.ExprMultiBool{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "is_used"}
	q.where = append(q.where, where)
	return q
}

func (q userItemDeleteSQL) HasExtension(v sql.NullBool, exprs ...sqlla.Operator) userItemDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}

	where := sqlla.ExprNullBool{Value: v, Op: op, Column: "has_extension"}
	q.where = append(q.where, where)
	return q
}


func (q userItemDeleteSQL) HasExtensionIn(vs ...sql.NullBool) userItemDeleteSQL {
	where := sqlla.ExprMultiNullBool{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "has_extension"}
	q.where = append(q.where, where)
	return q
}

func (q userItemDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM user_item"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func ( q userItemDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func ( q userItemDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s UserItem) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserItemSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s UserItem) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewUserItemSQL().Delete().ID(s.Id).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}

