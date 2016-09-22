package example

import (
	"database/sql"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mackee/go-sqlla"
)

func TestSelect(t *testing.T) {
	q := NewUserSQL().Select().Name("hoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT id, name FROM user WHERE name = ?;" {
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
	if query != "SELECT id, name FROM user WHERE name = ? ORDER BY id ASC LIMIT 100;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestUpdate(t *testing.T) {
	q := NewUserSQL().Update().SetName("barbar").WhereID(uint64(1))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "UPDATE user SET name = ? WHERE id = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"barbar", "1"}) {
		t.Error("unexpected args:", args)
	}
}

func TestInsert(t *testing.T) {
	q := NewUserSQL().Insert().ValueName("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "INSERT INTO user (name) VALUES(?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestDelete(t *testing.T) {
	q := NewUserSQL().Delete().Name("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "DELETE FROM user WHERE name = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
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
	err = row.Scan(&id, &name)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if id == uint64(0) {
		t.Error("empty id:", id)
	}

	user, err := NewUserSQL().Select().Name("hogehoge").Single(db)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if user.Name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if user.Id == uint64(0) {
		t.Error("empty id:", id)
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
