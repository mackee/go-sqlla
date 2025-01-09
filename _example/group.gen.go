// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla - DO NOT EDIT.
package example

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"database/sql"
	"time"

	"github.com/mackee/go-sqlla/v2"
)

type groupSQL struct {
	where sqlla.Where
}

func NewGroupSQL() groupSQL {
	q := groupSQL{}
	return q
}

var groupAllColumns = []string{
	"`id`", "`name`", "`leader_user_id`", "`sub_leader_user_id`", "`child_group_id`", "`created_at`", "`updated_at`",
}

type groupSelectSQL struct {
	groupSQL
	Columns     []string
	order       string
	limit       *uint64
	offset      *uint64
	tableAlias  string
	joinClauses []string

	additionalWhereClause     string
	additionalWhereClauseArgs []interface{}

	groupByColumns []string

	isForUpdate bool
}

func (q groupSQL) Select() groupSelectSQL {
	return groupSelectSQL{
		q,
		groupAllColumns,
		"",
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

func (q groupSelectSQL) Or(qs ...groupSelectSQL) groupSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q groupSelectSQL) Limit(l uint64) groupSelectSQL {
	q.limit = &l
	return q
}

func (q groupSelectSQL) Offset(o uint64) groupSelectSQL {
	q.offset = &o
	return q
}

func (q groupSelectSQL) ForUpdate() groupSelectSQL {
	q.isForUpdate = true
	return q
}

func (q groupSelectSQL) TableAlias(alias string) groupSelectSQL {
	q.tableAlias = "`" + alias + "`"
	return q
}

func (q groupSelectSQL) SetColumns(columns ...string) groupSelectSQL {
	q.Columns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(.`") {
			q.Columns = append(q.Columns, column)
		} else {
			q.Columns = append(q.Columns, "`"+column+"`")
		}
	}
	return q
}

func (q groupSelectSQL) JoinClause(clause string) groupSelectSQL {
	q.joinClauses = append(q.joinClauses, clause)
	return q
}

func (q groupSelectSQL) AdditionalWhereClause(clause string, args ...interface{}) groupSelectSQL {
	q.additionalWhereClause = clause
	q.additionalWhereClauseArgs = args
	return q
}

func (q groupSelectSQL) appendColumnPrefix(column string) string {
	if q.tableAlias == "" || strings.ContainsAny(column, "(.") {
		return column
	}
	return q.tableAlias + "." + column
}

func (q groupSelectSQL) GroupBy(columns ...string) groupSelectSQL {
	q.groupByColumns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(.`") {
			q.groupByColumns = append(q.groupByColumns, column)
		} else {
			q.groupByColumns = append(q.groupByColumns, "`"+column+"`")
		}
	}
	return q
}

func (q groupSelectSQL) ID(v GroupID, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) IDIn(vs ...GroupID) groupSelectSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) groupSelectSQL {
	v := GroupID(pk)
	return q.ID(v, exprs...)
}

func (q groupSelectSQL) OrderByID(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) Name(v string, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: q.appendColumnPrefix("`name`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) NameIn(vs ...string) groupSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`name`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderByName(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`name`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) LeaderUserID(v UserId, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: q.appendColumnPrefix("`leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) LeaderUserIDIn(vs ...UserId) groupSelectSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderByLeaderUserID(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`leader_user_id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) SubLeaderUserID(v int64, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: q.appendColumnPrefix("`sub_leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) SubLeaderUserIDIsNull() groupSelectSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: q.appendColumnPrefix("`sub_leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) SubLeaderUserIDIsNotNull() groupSelectSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: q.appendColumnPrefix("`sub_leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) SubLeaderUserIDIn(vs ...int64) groupSelectSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`sub_leader_user_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderBySubLeaderUserID(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`sub_leader_user_id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) ChildGroupID(v int64, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: q.appendColumnPrefix("`child_group_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) ChildGroupIDIsNull() groupSelectSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: q.appendColumnPrefix("`child_group_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) ChildGroupIDIsNotNull() groupSelectSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: q.appendColumnPrefix("`child_group_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) ChildGroupIDIn(vs ...int64) groupSelectSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`child_group_id`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderByChildGroupID(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`child_group_id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) CreatedAtIn(vs ...time.Time) groupSelectSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderByCreatedAt(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`created_at`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) UpdatedAt(v sql.NullTime, exprs ...sqlla.Operator) groupSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNullTime{Value: v, Op: op, Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) UpdatedAtIn(vs ...sql.NullTime) groupSelectSQL {
	where := sqlla.ExprMultiNullTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`updated_at`")}
	q.where = append(q.where, where)
	return q
}

func (q groupSelectSQL) OrderByUpdatedAt(order sqlla.Order) groupSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`updated_at`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q groupSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	tableName := "`group`"
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

func (s Group) Select() groupSelectSQL {
	return NewGroupSQL().Select().ID(s.ID)
}
func (q groupSelectSQL) Single(db sqlla.DB) (Group, error) {
	q.Columns = groupAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Group{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q groupSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (Group, error) {
	q.Columns = groupAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Group{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
	return q.Scan(row)
}

func (q groupSelectSQL) All(db sqlla.DB) ([]Group, error) {
	rs := make([]Group, 0, 10)
	q.Columns = groupAllColumns
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

func (q groupSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]Group, error) {
	rs := make([]Group, 0, 10)
	q.Columns = groupAllColumns
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

func (q groupSelectSQL) Scan(s sqlla.Scanner) (Group, error) {
	var row Group
	err := s.Scan(
		&row.ID,
		&row.Name,
		&row.LeaderUserID,
		&row.SubLeaderUserID,
		&row.ChildGroupID,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	return row, err
}

// IterContext returns iter.Seq2[Group, error] and closer.
//
// The returned Iter.Seq2 assembles and executes a query in the first iteration.
// Therefore, the first iteration may return an error in assembling or executing the query.
// Subsequent iterations read rows. Again, the read may return an error.
//
// closer is a function that closes the row reader object. Execution of this function is idempotent.
// Be sure to call it when you are done using iter.Seq2.
func (q groupSelectSQL) IterContext(ctx context.Context, db sqlla.DB) (func(func(Group, error) bool), func() error) {
	var rowClose func() error
	closer := func() error {
		if rowClose != nil {
			err := rowClose()
			rowClose = nil
			return err
		}
		return nil
	}

	q.Columns = groupAllColumns
	query, args, err := q.ToSql()
	return func(yield func(Group, error) bool) {
		if err != nil {
			var r Group
			yield(r, err)
			return
		}
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			var r Group
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

type groupUpdateSQL struct {
	groupSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q groupSQL) Update() groupUpdateSQL {
	return groupUpdateSQL{
		groupSQL: q,
		setMap:   sqlla.SetMap{},
	}
}

func (q groupUpdateSQL) SetID(v GroupID) groupUpdateSQL {
	q.setMap["`id`"] = v
	return q
}

func (q groupUpdateSQL) WhereID(v GroupID, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereIDIn(vs ...GroupID) groupUpdateSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetName(v string) groupUpdateSQL {
	q.setMap["`name`"] = v
	return q
}

func (q groupUpdateSQL) WhereName(v string, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereNameIn(vs ...string) groupUpdateSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetLeaderUserID(v UserId) groupUpdateSQL {
	q.setMap["`leader_user_id`"] = v
	return q
}

func (q groupUpdateSQL) WhereLeaderUserID(v UserId, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: "`leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereLeaderUserIDIn(vs ...UserId) groupUpdateSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetSubLeaderUserID(v int64) groupUpdateSQL {
	q.setMap["`sub_leader_user_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupUpdateSQL) SetSubLeaderUserIDToNull() groupUpdateSQL {
	q.setMap["`sub_leader_user_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupUpdateSQL) WhereSubLeaderUserID(v int64, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereSubLeaderUserIDIsNull() groupUpdateSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereSubLeaderUserIDIsNotNull() groupUpdateSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereSubLeaderUserIDIn(vs ...int64) groupUpdateSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetChildGroupID(v int64) groupUpdateSQL {
	q.setMap["`child_group_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupUpdateSQL) SetChildGroupIDToNull() groupUpdateSQL {
	q.setMap["`child_group_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupUpdateSQL) WhereChildGroupID(v int64, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereChildGroupIDIsNull() groupUpdateSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereChildGroupIDIsNotNull() groupUpdateSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereChildGroupIDIn(vs ...int64) groupUpdateSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetCreatedAt(v time.Time) groupUpdateSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q groupUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereCreatedAtIn(vs ...time.Time) groupUpdateSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) SetUpdatedAt(v sql.NullTime) groupUpdateSQL {
	q.setMap["`updated_at`"] = v
	return q
}

func (q groupUpdateSQL) WhereUpdatedAt(v sql.NullTime, exprs ...sqlla.Operator) groupUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNullTime{Value: v, Op: op, Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) WhereUpdatedAtIn(vs ...sql.NullTime) groupUpdateSQL {
	where := sqlla.ExprMultiNullTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Group{}
	if t, ok := s.(groupDefaultUpdateHooker); ok {
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

	query := "UPDATE `group` SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s Group) Update() groupUpdateSQL {
	return NewGroupSQL().Update().WhereID(s.ID)
}

func (q groupUpdateSQL) Exec(db sqlla.DB) ([]Group, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.groupSQL

	return qq.Select().All(db)
}

func (q groupUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]Group, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.groupSQL

	return qq.Select().AllContext(ctx, db)
}

type groupDefaultUpdateHooker interface {
	DefaultUpdateHook(groupUpdateSQL) (groupUpdateSQL, error)
}

type groupInsertSQL struct {
	groupSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q groupSQL) Insert() groupInsertSQL {
	return groupInsertSQL{
		groupSQL: q,
		setMap:   sqlla.SetMap{},
	}
}

func (q groupInsertSQL) ValueID(v GroupID) groupInsertSQL {
	q.setMap["`id`"] = v
	return q
}

func (q groupInsertSQL) ValueName(v string) groupInsertSQL {
	q.setMap["`name`"] = v
	return q
}

func (q groupInsertSQL) ValueLeaderUserID(v UserId) groupInsertSQL {
	q.setMap["`leader_user_id`"] = v
	return q
}

func (q groupInsertSQL) ValueSubLeaderUserID(v int64) groupInsertSQL {
	q.setMap["`sub_leader_user_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupInsertSQL) ValueSubLeaderUserIDIsNull() groupInsertSQL {
	q.setMap["`sub_leader_user_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupInsertSQL) ValueChildGroupID(v int64) groupInsertSQL {
	q.setMap["`child_group_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupInsertSQL) ValueChildGroupIDIsNull() groupInsertSQL {
	q.setMap["`child_group_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupInsertSQL) ValueCreatedAt(v time.Time) groupInsertSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q groupInsertSQL) ValueUpdatedAt(v sql.NullTime) groupInsertSQL {
	q.setMap["`updated_at`"] = v
	return q
}

func (q groupInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.groupInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q groupInsertSQL) groupInsertSQLToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Group{}
	if t, ok := s.(groupDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO `group` " + qs

	return query, vs, nil
}

func (q groupInsertSQL) OnDuplicateKeyUpdate() groupInsertOnDuplicateKeyUpdateSQL {
	return groupInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q groupInsertSQL) Exec(db sqlla.DB) (Group, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Group{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return Group{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Group{}, err
	}
	return NewGroupSQL().Select().PkColumn(id).Single(db)
}

func (q groupInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (Group, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Group{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return Group{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Group{}, err
	}
	return NewGroupSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q groupInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type groupDefaultInsertHooker interface {
	DefaultInsertHook(groupInsertSQL) (groupInsertSQL, error)
}

type groupInsertSQLToSqler interface {
	groupInsertSQLToSql() (string, []interface{}, error)
}

type groupInsertOnDuplicateKeyUpdateSQL struct {
	insertSQL               groupInsertSQLToSqler
	onDuplicateKeyUpdateMap sqlla.SetMap
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateID(v GroupID) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateID() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = sqlla.SetMapRawValue("VALUES(`id`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateName(v string) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateName(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateName() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`name`"] = sqlla.SetMapRawValue("VALUES(`name`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateLeaderUserID(v UserId) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`leader_user_id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateLeaderUserID(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`leader_user_id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateLeaderUserID() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`leader_user_id`"] = sqlla.SetMapRawValue("VALUES(`leader_user_id`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateSubLeaderUserID(v int64) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`sub_leader_user_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateSubLeaderUserIDToNull() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`sub_leader_user_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateSubLeaderUserID(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`sub_leader_user_id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateSubLeaderUserID() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`sub_leader_user_id`"] = sqlla.SetMapRawValue("VALUES(`sub_leader_user_id`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateChildGroupID(v int64) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`child_group_id`"] = sql.Null[int64]{V: v, Valid: true}
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateChildGroupIDToNull() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`child_group_id`"] = sql.Null[int64]{Valid: false}
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateChildGroupID(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`child_group_id`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateChildGroupID() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`child_group_id`"] = sqlla.SetMapRawValue("VALUES(`child_group_id`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateCreatedAt(v time.Time) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateCreatedAt() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = sqlla.SetMapRawValue("VALUES(`created_at`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateUpdatedAt(v sql.NullTime) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateUpdatedAt(v sqlla.SetMapRawValue) groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = v
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) SameOnUpdateUpdatedAt() groupInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`updated_at`"] = sqlla.SetMapRawValue("VALUES(`updated_at`)")
	return q
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Group{}
	if t, ok := s.(groupDefaultInsertOnDuplicateKeyUpdateHooker); ok {
		q, err = t.DefaultInsertOnDuplicateKeyUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}

	query, vs, err := q.insertSQL.groupInsertSQLToSql()
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

func (q groupInsertOnDuplicateKeyUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) (Group, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Group{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return Group{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Group{}, err
	}
	return NewGroupSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q groupInsertOnDuplicateKeyUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type groupDefaultInsertOnDuplicateKeyUpdateHooker interface {
	DefaultInsertOnDuplicateKeyUpdateHook(groupInsertOnDuplicateKeyUpdateSQL) (groupInsertOnDuplicateKeyUpdateSQL, error)
}

type groupBulkInsertSQL struct {
	insertSQLs []groupInsertSQL
}

func (q groupSQL) BulkInsert() *groupBulkInsertSQL {
	return &groupBulkInsertSQL{
		insertSQLs: []groupInsertSQL{},
	}
}

func (q *groupBulkInsertSQL) Append(iqs ...groupInsertSQL) {
	q.insertSQLs = append(q.insertSQLs, iqs...)
}

func (q *groupBulkInsertSQL) groupInsertSQLToSql() (string, []interface{}, error) {
	if len(q.insertSQLs) == 0 {
		return "", []interface{}{}, fmt.Errorf("sqlla: This groupBulkInsertSQL's InsertSQL was empty")
	}
	iqs := make([]groupInsertSQL, len(q.insertSQLs))
	copy(iqs, q.insertSQLs)

	var s interface{} = Group{}
	if t, ok := s.(groupDefaultInsertHooker); ok {
		for i, iq := range iqs {
			var err error
			iq, err = t.DefaultInsertHook(iq)
			if err != nil {
				return "", []interface{}{}, err
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
		return "", []interface{}{}, err
	}

	return "INSERT INTO `group` " + query, vs, nil
}

func (q *groupBulkInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.groupInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q *groupBulkInsertSQL) OnDuplicateKeyUpdate() groupInsertOnDuplicateKeyUpdateSQL {
	return groupInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q *groupBulkInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type groupDeleteSQL struct {
	groupSQL
}

func (q groupSQL) Delete() groupDeleteSQL {
	return groupDeleteSQL{
		q,
	}
}

func (q groupDeleteSQL) ID(v GroupID, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) IDIn(vs ...GroupID) groupDeleteSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) Name(v string, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) NameIn(vs ...string) groupDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`name`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) LeaderUserID(v UserId, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprUint64{Value: uint64(v), Op: op, Column: "`leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) LeaderUserIDIn(vs ...UserId) groupDeleteSQL {
	_vs := make([]uint64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, uint64(v))
	}
	where := sqlla.ExprMultiUint64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) SubLeaderUserID(v int64, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) SubLeaderUserIDIsNull() groupDeleteSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) SubLeaderUserIDIsNotNull() groupDeleteSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) SubLeaderUserIDIn(vs ...int64) groupDeleteSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`sub_leader_user_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) ChildGroupID(v int64, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{V: v, Valid: true}, Op: op, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) ChildGroupIDIsNull() groupDeleteSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpEqual, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) ChildGroupIDIsNotNull() groupDeleteSQL {
	where := sqlla.ExprNull[int64]{Value: sql.Null[int64]{Valid: false}, Op: sqlla.OpNot, Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) ChildGroupIDIn(vs ...int64) groupDeleteSQL {
	_vs := make([]sql.Null[int64], 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, sql.Null[int64]{V: v, Valid: true})
	}
	where := sqlla.ExprMultiValue[sql.Null[int64]]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`child_group_id`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprTime{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) CreatedAtIn(vs ...time.Time) groupDeleteSQL {
	where := sqlla.ExprMultiTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) UpdatedAt(v sql.NullTime, exprs ...sqlla.Operator) groupDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprNullTime{Value: v, Op: op, Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) UpdatedAtIn(vs ...sql.NullTime) groupDeleteSQL {
	where := sqlla.ExprMultiNullTime{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`updated_at`"}
	q.where = append(q.where, where)
	return q
}

func (q groupDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM `group`"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func (q groupDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (q groupDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s Group) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewGroupSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s Group) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewGroupSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
