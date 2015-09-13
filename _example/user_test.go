package example

import (
	"reflect"
	"testing"

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
