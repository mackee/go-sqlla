package main

import (
	"reflect"
	"testing"

	sq "github.com/lann/squirrel"
)

func BenchmarkSelect__Squirrel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := sq.Select("id", "name").From("user").Where(sq.Eq{"name": "hogehoge"})
		q.ToSql()
	}
}

func BenchmarkSelect__Sqlla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := example.NewUserSQL().Select().Name("hogehoge")
		q.ToSql()
	}
}

func TestSelect__Squirrel(t *testing.T) {
	q := sq.Select("id", "name").From("user").Where(sq.Eq{"name": "hogehoge"})
	query, args, err := q.ToSql()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if query != "SELECT id, name FROM user WHERE name = ?" {
		t.Fatal("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
		t.Fatal("unexpected args:", args)
	}
}

func TestSelect__Sqlla(t *testing.T) {
	q := example.NewUserSQL().Select().Name("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if query != "SELECT id, name FROM user WHERE name = ?;" {
		t.Fatal("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
		t.Fatal("unexpected args:", args)
	}
}
