package example_test

import (
	"context"
	"slices"
	"testing"

	example "github.com/mackee/go-sqlla/_example"
)

func TestPlugin__User__Count(t *testing.T) {
	ctx := context.Background()
	db := setupDB(t)

	for _, name := range []string{"hoge", "fuga", "piyo"} {
		if _, err := example.NewUserSQL().Insert().
			ValueName(name).
			ValueIconImage([]byte{}).
			ExecContextWithoutSelect(ctx, db); err != nil {
			t.Error("cannot insert row error:", err)
		}
	}

	allCount, err := example.NewUserSQL().Select().CountContext(ctx, db, "id")
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if allCount != 3 {
		t.Error("unexpected allCount:", allCount)
	}
	hogeCount, err := example.NewUserSQL().Select().Name("hoge").CountContext(ctx, db, "id")
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if hogeCount != 1 {
		t.Error("unexpected hogeCount:", hogeCount)
	}
	notFoundCount, err := example.NewUserSQL().Select().Name("notfound").CountContext(ctx, db, "id")
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if notFoundCount != 0 {
		t.Error("unexpected notFoundCount:", notFoundCount)
	}
}

func TestPlugin__User__List(t *testing.T) {
	ctx := context.Background()
	db := setupDB(t)

	for _, name := range []string{"hoge", "fuga", "piyo"} {
		if _, err := example.NewUserSQL().Insert().
			ValueName(name).
			ValueIconImage([]byte{}).
			ExecContextWithoutSelect(ctx, db); err != nil {
			t.Error("cannot insert row error:", err)
		}
	}

	_users, err := example.NewUserSQL().Select().AllContext(ctx, db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	users := make(example.Users, len(_users))
	for i, u := range _users {
		users[i] = &u
	}

	ids := users.Ids()
	slices.Sort(ids)
	if !slices.Equal(ids, []example.UserId{1, 2, 3}) {
		t.Error("unexpected ids:", ids)
	}

	names := users.Names()
	slices.Sort(names)
	if !slices.Equal(names, []string{"fuga", "hoge", "piyo"}) {
		t.Error("unexpected names:", names)
	}

	userIdMap := users.AssociateByIds()
	expectedUserNameMap := map[example.UserId]string{
		1: "hoge",
		2: "fuga",
		3: "piyo",
	}
	for id, u := range userIdMap {
		if id != u.Id {
			t.Error("unexpected id:", id)
		}
		if expectedUserNameMap[id] != u.Name {
			t.Error("unexpected name:", u.Name)
		}
	}

	userNameMap := users.GroupByNames()
	expectedUserIdMap := map[string]example.UserId{
		"hoge": 1,
		"fuga": 2,
		"piyo": 3,
	}
	for name, us := range userNameMap {
		if !slices.Equal(us.Names(), []string{name}) {
			t.Error("unexpected names:", us.Names())
		}
		if !slices.Equal(us.Ids(), []example.UserId{expectedUserIdMap[name]}) {
			t.Error("unexpected names:", us.Ids())
		}
	}
}
