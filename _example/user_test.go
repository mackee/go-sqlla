package example

import (
	"database/sql"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mackee/go-sqlla"
	_ "github.com/mattn/go-sqlite3"
)

var columns = "`id`, `name`, `age`, `rate`, `created_at`, `updated_at`"

func TestSelect(t *testing.T) {
	q := NewUserSQL().Select().Name("hoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+columns+" FROM user WHERE `name` = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__OrderByAndLimit(t *testing.T) {
	q := NewUserSQL().Select().Name("hoge").OrderByID(sqlla.Asc).Limit(uint64(100))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+columns+" FROM user WHERE `name` = ? ORDER BY `id` ASC LIMIT 100;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__InOperator(t *testing.T) {
	q := NewUserSQL().Select().IDIn(1, 2, 3, 4, 5)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+columns+" FROM user WHERE `id` IN(?,?,?,?,?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{uint64(1), uint64(2), uint64(3), uint64(4), uint64(5)}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__NullInt64(t *testing.T) {
	q := NewUserSQL().Select().Age(sql.NullInt64{})
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+columns+" FROM user WHERE `age` IS NULL;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__ForUpdate(t *testing.T) {
	q := NewUserSQL().Select().ID(uint64(1)).ForUpdate()
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+columns+" FROM user WHERE `id` = ? FOR UPDATE;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"1"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__Or(t *testing.T) {
	q := NewUserSQL().Select().Or(
		NewUserSQL().Select().ID(uint64(1)),
		NewUserSQL().Select().ID(uint64(2)),
	)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT " + columns + " FROM user WHERE (( `id` = ? ) OR ( `id` = ? ));"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{"1", "2"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__OrNull(t *testing.T) {
	now := time.Now()
	nt := mysql.NullTime{Time: now, Valid: true}
	q := NewUserItemSQL().Select().
		IDIn(uint64(1), uint64(2)).
		Or(
			NewUserItemSQL().Select().UsedAt(mysql.NullTime{}, sqlla.OpIs),
			NewUserItemSQL().Select().UsedAt(nt, sqlla.OpLess),
		)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT `id`, `user_id`, `item_id`, `is_used`, `has_extension`, `used_at` FROM user_item WHERE `id` IN(?,?) AND (( `used_at` IS NULL ) OR ( `used_at` < ? ));"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{uint64(1), uint64(2), nt}) {
		t.Error("unexpected args:", args)
	}
}

func TestUpdate(t *testing.T) {
	q := NewUserSQL().Update().SetName("barbar").WhereID(uint64(1))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	switch query {
	case "UPDATE user SET `name` = ?, `updated_at` = ? WHERE `id` = ?;":
		if !reflect.DeepEqual(args[0], "barbar") {
			t.Error("unexpected args:", args)
		}
		if !reflect.DeepEqual(args[2], "1") {
			t.Error("unexpected args:", args)
		}
	case "UPDATE user SET `updated_at` = ?, `name` = ? WHERE `id` = ?;":
		if !reflect.DeepEqual(args[2], "1") {
			t.Error("unexpected args:", args)
		}
		if !reflect.DeepEqual(args[1], "barbar") {
			t.Error("unexpected args:", args)
		}
	default:
		t.Error("unexpected query:", query)
	}
}

func TestInsert(t *testing.T) {
	q := NewUserSQL().Insert().ValueName("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	switch query {
	case "INSERT INTO user (`name`,`created_at`) VALUES(?,?);":
		if !reflect.DeepEqual(args[0], "hogehoge") {
			t.Error("unexpected args:", args)
		}
	case "INSERT INTO user (`created_at`,`name`) VALUES(?,?);":
		if !reflect.DeepEqual(args[1], "hogehoge") {
			t.Error("unexpected args:", args)
		}
	default:
		t.Error("unexpected query:", query)
	}
}

func TestDelete(t *testing.T) {
	q := NewUserSQL().Delete().Name("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "DELETE FROM user WHERE `name` = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestDelete__In(t *testing.T) {
	q := NewUserSQL().Delete().NameIn("hogehoge", "fugafuga")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "DELETE FROM user WHERE `name` IN(?,?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge", "fugafuga"}) {
		t.Error("unexpected args:", args)
	}
}

func setupDB(t *testing.T) *sql.DB {
	dbFile, err := ioutil.TempFile("", "sqlla_test")
	if err != nil {
		t.Fatal("cannot create tempfile error:", err)
	}
	db, err := sql.Open("sqlite3", dbFile.Name())
	if err != nil {
		t.Fatal("cannot open database error:", err)
	}

	schemaFile, err := os.Open("./sqlite3.sql")
	if err != nil {
		t.Fatal("cannot open schema file error:", err)
	}

	b, err := ioutil.ReadAll(schemaFile)
	if err != nil {
		t.Fatal("cannot read schema file error:", err)
	}

	stmts := strings.Split(string(b), ";")
	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			t.Fatal("cannot load schema error:", err)
		}
	}
	return db
}

func TestCRUD__WithSqlite3(t *testing.T) {
	db := setupDB(t)

	query, args, err := NewUserSQL().Insert().ValueName("hogehoge").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}

	query, args, err = NewUserSQL().Select().Name("hogehoge").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	row := db.QueryRow(query, args...)
	var id uint64
	var name string
	var age sql.NullInt64
	var rate float64
	var createdAt time.Time
	var updatedAt mysql.NullTime
	err = row.Scan(&id, &name, &age, &rate, &createdAt, &updatedAt)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if id == uint64(0) {
		t.Error("empty id:", id)
	}

	query, args, err = NewUserSQL().Select().IDIn(uint64(1)).ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	row = db.QueryRow(query, args...)
	var rescanID uint64
	var rescanName string
	var rescanAge sql.NullInt64
	var rescanRate float64
	var rescanCreatedAt time.Time
	var rescanUpdatedAt mysql.NullTime
	err = row.Scan(&rescanID, &rescanName, &rescanAge, &rescanRate, &rescanCreatedAt, &rescanUpdatedAt)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if rescanID != id {
		t.Error("unmatched id:", rescanID)
	}

	query, args, err = NewUserSQL().Update().WhereID(id).SetName("barbar").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		t.Error("cannot update row error:", err)
	}
	if rows, _ := result.RowsAffected(); rows != 1 {
		t.Error("unexpected row affected:", rows)
	}

	query, args, err = NewUserSQL().Delete().Name("barbar").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	result, err = db.Exec(query, args...)
	if err != nil {
		t.Error("cannot delete row error:", err)
	}
	if rows, _ := result.RowsAffected(); rows != 1 {
		t.Error("unexpected row affected:", rows)
	}
}

func TestORM__WithSqlite3(t *testing.T) {
	db := setupDB(t)

	insertedRow, err := NewUserSQL().Insert().ValueName("hogehoge").Exec(db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}
	if insertedRow.Id == uint64(0) {
		t.Error("empty id:", insertedRow.Id)
	}
	if insertedRow.Name != "hogehoge" {
		t.Error("unexpected name:", insertedRow.Name)
	}

	singleRow, err := NewUserSQL().Select().ID(insertedRow.Id).Single(db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if singleRow.Id == uint64(0) {
		t.Error("empty id:", singleRow.Id)
	}
	if singleRow.Name != "hogehoge" {
		t.Error("unexpected name:", singleRow.Name)
	}

	_, err = NewUserSQL().Insert().ValueName("fugafuga").Exec(db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}

	rows, err := NewUserSQL().Select().All(db)
	if len(rows) != 2 {
		t.Error("missing rows error:", len(rows))
	}

	for _, row := range rows {
		if row.Id == uint64(0) {
			t.Error("empty id:", row.Id)
		}
		if row.Name != "hogehoge" && row.Name != "fugafuga" {
			t.Error("unexpected name:", row.Name)
		}
	}

	targetRow := rows[0]
	results, err := targetRow.Update().SetName("barbar").Exec(db)
	if err != nil {
		t.Error("cannnot update row error", err)
	}
	if len(results) != 1 {
		t.Error("unexpected rows results:", len(results))
	}
	result := results[0]
	if result.Id != targetRow.Id {
		t.Errorf("result.Id is not targetRow.Id: %d vs %d", result.Id, targetRow.Id)
	}
	if result.Name != "barbar" {
		t.Errorf("result.Name is not replaced to \"barbar\": %s", result.Name)
	}

	deletedResult, err := targetRow.Delete(db)
	if err != nil {
		t.Error("cannnot delete row error", err)
	}
	if affected, _ := deletedResult.RowsAffected(); affected != int64(1) {
		t.Error("unexpected rows affected:", affected)
	}

	_, err = targetRow.Select().Single(db)
	if err != sql.ErrNoRows {
		t.Error("not deleted rows")
	}
}
