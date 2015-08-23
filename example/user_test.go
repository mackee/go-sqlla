package example

import (
	"reflect"
	"testing"

	"github.com/mackee/go-sqlla"
)

func TestSelect(t *testing.T) {
	q := NewUserSQL().Select().Name("hoge")
	query, args, err := q.ToSelectSql()
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
	q := NewUserSQL().Select().Name("hoge").OrderByID(sqlla.Asc)
	query, args, err := q.ToSelectSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT id, name FROM user WHERE name = ? ORDER BY id ASC;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}
