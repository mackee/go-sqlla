package example_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	example "github.com/mackee/go-sqlla/_example"
)

func TestPlugin__Group__Table(t *testing.T) {
	ctx := context.Background()
	db := setupDB(t)

	table := example.NewGroupTable()

	now := time.Now()
	inputs := make([]example.GroupTableCreateInput, 4)
	for i, name := range []string{"hoge", "fuga", "piyo", "barr"} {
		inputs[i] = example.GroupTableCreateInput{
			Name:         name,
			LeaderUserID: 1,
			CreatedAt:    now,
		}
		if i%2 == 0 {
			inputs[i].SubLeaderUserID = sql.Null[int64]{
				Valid: true,
				V:     2,
			}
		}
	}
	if err := table.CreateMulti(ctx, db, inputs); err != nil {
		t.Error("cannot create rows error:", err)
	}

	row1, err := table.GetByIDAndLeaderUserID(ctx, db, 1, 1)
	if err != nil {
		t.Error("cannot get row error:", err)
	}
	if row1.SubLeaderUserID.V != 2 {
		t.Error("unexpected row1.SubLeaderUserID.V:", row1.SubLeaderUserID.V)
	}

	rows, err := table.ListByLeaderUserIDAndSubLeaderUserID(ctx, db, 1, 2)
	if err != nil {
		t.Error("cannot list rows error:", err)
	}
	if len(rows) != 2 {
		t.Error("unexpected len(rows):", len(rows))
	}

	allRows, closer := example.NewGroupSQL().Select().IterContext(ctx, db)
	var ids []example.GroupID
	func() {
		defer closer()
		for row, err := range allRows {
			if err != nil {
				t.Error("cannot iterate rows error:", err)
			}
			if row.ID == 0 {
				t.Error("unexpected row.ID:", row.ID)
			}
			ids = append(ids, row.ID)
		}
	}()
	expectedRows, err := table.ListByIDS(ctx, db, ids)
	if err != nil {
		t.Error("cannot list rows by ids error:", err)
	}
	if len(expectedRows) != len(ids) {
		t.Error("unexpected len(expectedRows):", len(expectedRows))
	}
}
