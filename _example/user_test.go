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

func TestSelect__InOperator(t *testing.T) {
	q := NewUserSQL().Select().IDIn(1, 2, 3, 4, 5)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT id, name FROM user WHERE id IN(?,?,?,?,?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{uint64(1), uint64(2), uint64(3), uint64(4), uint64(5)}) {
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

func TestCRUD__WithSqlite3(t *testing.T) {
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

	query, args, err = NewUserSQL().Select().IDIn(uint64(1)).ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	row = db.QueryRow(query, args...)
	var rescanId uint64
	var rescanName string
	err = row.Scan(&rescanId, &rescanName)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if rescanId != id {
		t.Error("unmatched id:", rescanId)
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
