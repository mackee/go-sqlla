package example_test

import (
	"context"
	"database/sql"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	example "github.com/mackee/go-sqlla/_example"
	"github.com/mackee/go-sqlla/v2"
	_ "github.com/mattn/go-sqlite3"
)

var userAllColumns = strings.Join(example.UserAllColumns, ", ")

func TestQuery(t *testing.T) {
	sampleDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	ignoreDate := time.Date(2019, 12, 31, 23, 59, 59, 0, time.UTC)
	type toSqler interface {
		ToSql() (string, []any, error)
	}
	tcs := []struct {
		name     string
		query    toSqler
		expected string
		vs       []any
	}{
		{
			name:     "select",
			query:    example.NewUserSQL().Select().Name("hoge"),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `name` = ?;",
			vs:       []any{"hoge"},
		},
		{
			name:     "select with order by and limit",
			query:    example.NewUserSQL().Select().Name("hoge").OrderByID(sqlla.Asc).Limit(100),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `name` = ? ORDER BY `id` ASC LIMIT 100;",
			vs:       []any{"hoge"},
		},
		{
			name:     "select with in operator",
			query:    example.NewUserSQL().Select().IDIn(1, 2, 3, 4, 5),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `id` IN(?,?,?,?,?);",
			vs:       []any{uint64(1), uint64(2), uint64(3), uint64(4), uint64(5)},
		},
		{
			name:     "select with null int64",
			query:    example.NewUserSQL().Select().Age(sql.NullInt64{}),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `age` IS NULL;",
			vs:       []any{},
		},
		{
			name:     "select with not null int64",
			query:    example.NewUserSQL().Select().Age(sql.NullInt64{}, sqlla.OpNot),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `age` IS NOT NULL;",
			vs:       []any{},
		},
		{
			name:     "select with for update",
			query:    example.NewUserSQL().Select().ID(1).ForUpdate(),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `id` = ? FOR UPDATE;",
			vs:       []any{uint64(1)},
		},
		{
			name: "select or",
			query: example.NewUserSQL().Select().Or(
				example.NewUserSQL().Select().ID(1),
				example.NewUserSQL().Select().ID(2),
			),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE (( `id` = ? ) OR ( `id` = ? ));",
			vs:       []any{uint64(1), uint64(2)},
		},
		{
			name: "select or null",
			query: example.NewUserItemSQL().Select().
				IDIn(1, 2).
				Or(
					example.NewUserItemSQL().Select().UsedAt(sql.NullTime{}, sqlla.OpIs),
					example.NewUserItemSQL().Select().UsedAt(sql.NullTime{Time: sampleDate, Valid: true}, sqlla.OpLess),
				),
			expected: "SELECT `id`, `user_id`, `item_id`, `is_used`, `has_extension`, `used_at` FROM `user_item` WHERE `id` IN(?,?) AND (( `used_at` IS NULL ) OR ( `used_at` < ? ));",
			vs:       []any{uint64(1), uint64(2), sampleDate},
		},
		{
			name: "select join clause and table alias",
			query: example.NewUserSQL().Select().
				SetColumns(append(example.UserAllColumns, "ui.item_id", "ui.is_used")...).
				TableAlias("u").
				JoinClause("INNER JOIN user_item AS ui ON u.id = ui.user_id").
				Name("hogehoge").
				AdditionalWhereClause("AND ui.item_id IN (?,?,?)", 1, 2, 3).
				OrderByID(sqlla.Desc),
			expected: "SELECT `u`.`id`, `u`.`name`, `u`.`age`, `u`.`rate`, `u`.`icon_image`, `u`.`created_at`, `u`.`updated_at`, ui.item_id, ui.is_used FROM `user` AS `u` INNER JOIN user_item AS ui ON u.id = ui.user_id WHERE `u`.`name` = ? AND ui.item_id IN (?,?,?) ORDER BY `u`.`id` DESC;",
			vs:       []any{"hogehoge", int(1), int(2), int(3)},
		},
		{
			name:     "select set column",
			query:    example.NewUserSQL().Select().SetColumns("rate", "COUNT(u.id)").TableAlias("u").OrderByRate(sqlla.Desc).GroupBy("rate"),
			expected: "SELECT `u`.`rate`, COUNT(u.id) FROM `user` AS `u` GROUP BY `u`.`rate` ORDER BY `u`.`rate` DESC;",
			vs:       nil,
		},
		{
			name:     "select group by dotted column",
			query:    example.NewUserSQL().Select().SetColumns("rate", "COUNT(u.id)").TableAlias("u").OrderByRate(sqlla.Desc).GroupBy("u.rate"),
			expected: "SELECT `u`.`rate`, COUNT(u.id) FROM `user` AS `u` GROUP BY u.rate ORDER BY `u`.`rate` DESC;",
			vs:       nil,
		},
		{
			name:     "select like operator",
			query:    example.NewUserSQL().Select().Name("%foobar%", sqlla.OpLike),
			expected: "SELECT " + userAllColumns + " FROM `user` WHERE `name` LIKE ?;",
			vs:       []any{string("%foobar%")},
		},
		{
			name:     "update",
			query:    example.NewUserSQL().Update().SetName("barbar").WhereID(example.UserId(1)),
			expected: "UPDATE `user` SET `name` = ?, `updated_at` = ? WHERE `id` = ?;",
			vs:       []any{"barbar", sql.Null[time.Time]{V: ignoreDate, Valid: true}, uint64(1)},
		},
		{
			name:     "update in operator",
			query:    example.NewUserSQL().Update().SetRate(42).WhereIDIn(example.UserId(1), example.UserId(2), example.UserId(3)),
			expected: "UPDATE `user` SET `rate` = ?, `updated_at` = ? WHERE `id` IN(?,?,?);",
			vs:       []any{float64(42), sql.Null[time.Time]{V: ignoreDate, Valid: true}, uint64(1), uint64(2), uint64(3)},
		},
		{
			name:     "insert",
			query:    example.NewUserSQL().Insert().ValueName("hogehoge"),
			expected: "INSERT INTO `user` (`created_at`,`name`) VALUES (?,?);",
			vs:       []any{ignoreDate, "hogehoge"},
		},
		{
			name: "insert on duplicate key update",
			query: example.NewUserSQL().Insert().
				ValueID(1).
				ValueName("hogehoge").
				ValueUpdatedAt(mysql.NullTime{
					Valid: true,
					Time:  sampleDate,
				}).
				OnDuplicateKeyUpdate().
				ValueOnUpdateAge(sql.NullInt64{
					Valid: true,
					Int64: 17,
				}),
			expected: "INSERT INTO `user` (`created_at`,`id`,`name`,`updated_at`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `age` = ?, `updated_at` = VALUES(`updated_at`);",
			vs:       []any{ignoreDate, uint64(1), "hogehoge", sql.Null[time.Time]{V: sampleDate, Valid: true}, sql.Null[int64]{V: 17, Valid: true}},
		},
		{
			name: "bulk insert",
			query: func() toSqler {
				items := example.NewUserItemSQL().BulkInsert()
				for i := 1; i <= 10; i++ {
					q := example.NewUserItemSQL().Insert().
						ValueUserID(42).
						ValueItemID(strconv.Itoa(i))
					items.Append(q)
				}
				return items
			}(),
			expected: "INSERT INTO `user_item` (`item_id`,`user_id`) VALUES (?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?);",
			vs:       []any{"1", uint64(42), "2", uint64(42), "3", uint64(42), "4", uint64(42), "5", uint64(42), "6", uint64(42), "7", uint64(42), "8", uint64(42), "9", uint64(42), "10", uint64(42)},
		},
		{
			name: "bulk insert with on duplicate key update",
			query: func() toSqler {
				items := example.NewUserItemSQL().BulkInsert()
				items.Append(
					example.NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("1").ValueIsUsed(true),
					example.NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("2").ValueIsUsed(true),
				)

				return items.OnDuplicateKeyUpdate().
					SameOnUpdateIsUsed().
					ValueOnUpdateUsedAt(sql.NullTime{
						Time:  sampleDate,
						Valid: true,
					})
			}(),
			expected: "INSERT INTO `user_item` (`is_used`,`item_id`,`user_id`) VALUES (?,?,?),(?,?,?) ON DUPLICATE KEY UPDATE `is_used` = VALUES(`is_used`), `used_at` = ?;",
			vs:       []any{true, "1", uint64(42), true, "2", uint64(42), sql.Null[time.Time]{V: sampleDate, Valid: true}},
		},
		{
			name:     "delete",
			query:    example.NewUserSQL().Delete().Name("hogehoge"),
			expected: "DELETE FROM `user` WHERE `name` = ?;",
			vs:       []any{"hogehoge"},
		},
		{
			name:     "delete in operator",
			query:    example.NewUserSQL().Delete().NameIn("hogehoge", "fugafuga"),
			expected: "DELETE FROM `user` WHERE `name` IN(?,?);",
			vs:       []any{"hogehoge", "fugafuga"},
		},
	}
	opts := cmp.Options{
		cmp.FilterValues(
			func(x, y time.Time) bool {
				return x == ignoreDate || y == ignoreDate
			},
			cmp.Ignore(),
		),
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			sql, vs, err := tc.query.ToSql()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if sql != tc.expected {
				t.Errorf("expected: \n%s\n, but got: \n%s", tc.expected, sql)
			}
			if len(vs) != len(tc.vs) {
				t.Errorf("expected: %v, but got: %v", tc.vs, vs)
			}
			if diff := cmp.Diff(vs, tc.vs, opts...); diff != "" {
				t.Errorf("has diff: %s", diff)
			}
		})
	}
}

func setupDB(t *testing.T) *sql.DB {
	dbFile, err := os.CreateTemp("", "sqlla_test")
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

	b, err := io.ReadAll(schemaFile)
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

	query, args, err := example.NewUserSQL().Insert().ValueName("hogehoge").ValueIconImage([]byte{}).ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}

	query, args, err = example.NewUserSQL().Select().Name("hogehoge").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	row := db.QueryRow(query, args...)
	var id uint64
	var name string
	var age sql.NullInt64
	var rate float64
	var iconImage []byte
	var createdAt time.Time
	var updatedAt mysql.NullTime
	err = row.Scan(&id, &name, &age, &rate, &iconImage, &createdAt, &updatedAt)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if id == uint64(0) {
		t.Error("empty id:", id)
	}

	query, args, err = example.NewUserSQL().Select().IDIn(example.UserId(1)).ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	row = db.QueryRow(query, args...)
	var rescanID uint64
	var rescanName string
	var rescanAge sql.NullInt64
	var rescanRate float64
	var rescanIconImage []byte
	var rescanCreatedAt time.Time
	var rescanUpdatedAt mysql.NullTime
	err = row.Scan(&rescanID, &rescanName, &rescanAge, &rescanRate, &rescanIconImage, &rescanCreatedAt, &rescanUpdatedAt)
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if name != "hogehoge" {
		t.Error("unexpected name:", name)
	}
	if rescanID != id {
		t.Error("unmatched id:", rescanID)
	}

	query, args, err = example.NewUserSQL().Update().WhereID(example.UserId(id)).SetName("barbar").ToSql()
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

	query, args, err = example.NewUserSQL().Delete().Name("barbar").ToSql()
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

	insertedRow, err := example.NewUserSQL().Insert().ValueName("hogehoge").ValueIconImage([]byte{}).Exec(db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}
	if insertedRow.Id == example.UserId(0) {
		t.Error("empty id:", insertedRow.Id)
	}
	if insertedRow.Name != "hogehoge" {
		t.Error("unexpected name:", insertedRow.Name)
	}

	singleRow, err := example.NewUserSQL().Select().ID(insertedRow.Id).Single(db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if singleRow.Id == example.UserId(0) {
		t.Error("empty id:", singleRow.Id)
	}
	if singleRow.Name != "hogehoge" {
		t.Error("unexpected name:", singleRow.Name)
	}

	_, err = example.NewUserSQL().Insert().ValueName("fugafuga").ValueIconImage([]byte{}).Exec(db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}

	rows, err := example.NewUserSQL().Select().All(db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if len(rows) != 2 {
		t.Error("missing rows error:", len(rows))
	}

	for _, row := range rows {
		if row.Id == example.UserId(0) {
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

func TestORM__WithSqlite3__Binary(t *testing.T) {
	db := setupDB(t)
	binary := []byte("binary")

	insertedRow, err := example.NewUserSQL().Insert().
		ValueName("hogehoge").
		ValueIconImage(binary).
		Exec(db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}
	if !reflect.DeepEqual(insertedRow.IconImage, binary) {
		t.Error("unexpected IconImage:", insertedRow.IconImage)
	}

	singleRow, err := example.NewUserSQL().Select().ID(insertedRow.Id).Single(db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if !reflect.DeepEqual(singleRow.IconImage, binary) {
		t.Error("unexpected IconImage:", singleRow.IconImage)
	}

	updatedBinary := []byte("updated")
	results, err := singleRow.Update().SetIconImage(updatedBinary).Exec(db)
	if err != nil {
		t.Error("cannnot update row error", err)
	}
	if len(results) != 1 {
		t.Error("unexpected rows results:", len(results))
	}
	result := results[0]
	if !reflect.DeepEqual(result.IconImage, updatedBinary) {
		t.Errorf("result.IconImage is not replaced to \"updated\": %s", result.IconImage)
	}
}

func TestORM__WithSqlite3__NullBinary(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	db := setupDB(t)

	_, err := example.NewUserExternalSQL().Insert().
		ValueID(42).
		ValueUserID(4242).
		ValueIconImage(nil).
		ValueCreatedAt(now).
		ValueUpdatedAt(now).
		ExecContextWithoutSelect(ctx, db)
	if err != nil {
		t.Error("cannot insert row error:", err)
	}

	singleRow, err := example.NewUserExternalSQL().Select().ID(42).SingleContext(ctx, db)
	if err != nil {
		t.Error("cannot select row error:", err)
	}
	if singleRow.IconImage != nil {
		t.Error("unexpected IconImage:", singleRow.IconImage)
	}

	updatedBinary := []byte("updated")
	results, err := singleRow.Update().SetIconImage(updatedBinary).ExecContext(ctx, db)
	if err != nil {
		t.Error("cannnot update row error", err)
	}
	if len(results) != 1 {
		t.Error("unexpected rows results:", len(results))
	}
	result := results[0]
	if !reflect.DeepEqual(result.IconImage, updatedBinary) {
		t.Errorf("result.IconImage is not replaced to \"updated\": %s", result.IconImage)
	}
}
