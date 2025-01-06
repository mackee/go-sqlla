package example_test

import (
	"context"
	"database/sql"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	example "github.com/mackee/go-sqlla/_example"
	"github.com/mackee/go-sqlla/v2"
	_ "github.com/mattn/go-sqlite3"
)

var userAllColumns = strings.Join(example.UserAllColumns, ", ")

func TestSelect(t *testing.T) {
	q := example.NewUserSQL().Select().Name("hoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `name` = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__OrderByAndLimit(t *testing.T) {
	q := example.NewUserSQL().Select().Name("hoge").OrderByID(sqlla.Asc).Limit(uint64(100))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `name` = ? ORDER BY `id` ASC LIMIT 100;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__InOperator(t *testing.T) {
	q := example.NewUserSQL().Select().IDIn(1, 2, 3, 4, 5)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `id` IN(?,?,?,?,?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{uint64(1), uint64(2), uint64(3), uint64(4), uint64(5)}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__NullInt64(t *testing.T) {
	q := example.NewUserSQL().Select().Age(sql.NullInt64{})
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `age` IS NULL;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__NullInt64__IsNotNull(t *testing.T) {
	q := example.NewUserSQL().Select().Age(sql.NullInt64{}, sqlla.OpNot)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `age` IS NOT NULL;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__ForUpdate(t *testing.T) {
	q := example.NewUserSQL().Select().ID(example.UserId(1)).ForUpdate()
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT "+userAllColumns+" FROM `user` WHERE `id` = ? FOR UPDATE;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"1"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__Or(t *testing.T) {
	q := example.NewUserSQL().Select().Or(
		example.NewUserSQL().Select().ID(example.UserId(1)),
		example.NewUserSQL().Select().ID(example.UserId(2)),
	)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT " + userAllColumns + " FROM `user` WHERE (( `id` = ? ) OR ( `id` = ? ));"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{"1", "2"}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__OrNull(t *testing.T) {
	now := time.Now()
	nt := sql.NullTime{Time: now, Valid: true}
	q := example.NewUserItemSQL().Select().
		IDIn(uint64(1), uint64(2)).
		Or(
			example.NewUserItemSQL().Select().UsedAt(sql.NullTime{}, sqlla.OpIs),
			example.NewUserItemSQL().Select().UsedAt(nt, sqlla.OpLess),
		)
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT `id`, `user_id`, `item_id`, `is_used`, `has_extension`, `used_at` FROM `user_item` WHERE `id` IN(?,?) AND (( `used_at` IS NULL ) OR ( `used_at` < ? ));"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{uint64(1), uint64(2), nt}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__JoinClausesAndTableAlias(t *testing.T) {
	query, args, err := example.NewUserSQL().Select().
		SetColumns(append(example.UserAllColumns, "ui.item_id", "ui.is_used")...).
		TableAlias("u").
		JoinClause("INNER JOIN user_item AS ui ON u.id = ui.user_id").
		Name("hogehoge").
		AdditionalWhereClause("AND ui.item_id IN (?,?,?)", 1, 2, 3).
		OrderByID(sqlla.Desc).
		ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT `u`.`id`, `u`.`name`, `u`.`age`, `u`.`rate`, `u`.`icon_image`, `u`.`created_at`, `u`.`updated_at`, ui.item_id, ui.is_used FROM `user` AS `u` INNER JOIN user_item AS ui ON u.id = ui.user_id WHERE `u`.`name` = ? AND ui.item_id IN (?,?,?) ORDER BY `u`.`id` DESC;"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge", int(1), int(2), int(3)}) {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__SetColumn(t *testing.T) {
	query, args, err := example.NewUserSQL().Select().
		SetColumns("rate", "COUNT(u.id)").
		TableAlias("u").
		OrderByRate(sqlla.Desc).
		GroupBy("rate").
		ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT `u`.`rate`, COUNT(u.id) FROM `user` AS `u` GROUP BY `u`.`rate` ORDER BY `u`.`rate` DESC;"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if len(args) != 0 {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__GroupByDottedColumn(t *testing.T) {
	query, args, err := example.NewUserSQL().Select().
		SetColumns("rate", "COUNT(u.id)").
		TableAlias("u").
		OrderByRate(sqlla.Desc).
		GroupBy("u.rate").
		ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT `u`.`rate`, COUNT(u.id) FROM `user` AS `u` GROUP BY u.rate ORDER BY `u`.`rate` DESC;"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if len(args) != 0 {
		t.Error("unexpected args:", args)
	}
}

func TestSelect__LikeOperator(t *testing.T) {
	query, args, err := example.NewUserSQL().Select().
		Name("%foobar%", sqlla.OpLike).
		ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expectedQuery := "SELECT " + userAllColumns + " FROM `user` WHERE `name` LIKE ?;"
	if query != expectedQuery {
		t.Error("unexpected query:", query, expectedQuery)
	}
	if !reflect.DeepEqual(args, []interface{}{string("%foobar%")}) {
		t.Error("unexpected args:", args)
	}
}

func TestUpdate(t *testing.T) {
	q := example.NewUserSQL().Update().SetName("barbar").WhereID(example.UserId(1))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	switch query {
	case "UPDATE `user` SET `name` = ?, `updated_at` = ? WHERE `id` = ?;":
		if !reflect.DeepEqual(args[0], "barbar") {
			t.Error("unexpected args:", args)
		}
		if !reflect.DeepEqual(args[2], "1") {
			t.Error("unexpected args:", args)
		}
	case "UPDATE `user` SET `updated_at` = ?, `name` = ? WHERE `id` = ?;":
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

func TestUpdate__InOperator(t *testing.T) {
	q := example.NewUserSQL().Update().SetRate(42).WhereIDIn(example.UserId(1), example.UserId(2), example.UserId(3))
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	switch query {
	case "UPDATE `user` SET `rate` = ?, `updated_at` = ? WHERE `id` IN(?,?,?);":
		if !reflect.DeepEqual(args[0], float64(42)) {
			t.Errorf("unexpected args: %+v", args[0])
		}
		for i, v := range []uint64{1, 2, 3} {
			if !reflect.DeepEqual(args[2+i], v) {
				t.Errorf("unexpected args: i=%d, %+v", i, args[2:])
			}
		}
	case "UPDATE `user` SET `updated_at` = ?, `rate` = ? WHERE `id` IN(?,?,?);":
		if !reflect.DeepEqual(args[1], float64(42)) {
			t.Errorf("unexpected args: %+v", args[1])
		}
		for i, v := range []uint64{1, 2, 3} {
			if !reflect.DeepEqual(args[2+i], v) {
				t.Errorf("unexpected args: i=%d, %+v", i, args[2:])
			}
		}
	default:
		t.Error("unexpected query:", query)
	}
}

func TestInsert(t *testing.T) {
	q := example.NewUserSQL().Insert().ValueName("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expected := "INSERT INTO `user` (`created_at`,`name`) VALUES(?,?);"
	if query != expected {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args[1], "hogehoge") {
		t.Error("unexpected args:", args)
	}
}

func TestInsertOnDuplicateKeyUpdate(t *testing.T) {
	now := time.Now()
	q := example.NewUserSQL().Insert().
		ValueID(1).
		ValueName("hogehoge").
		ValueUpdatedAt(mysql.NullTime{
			Valid: true,
			Time:  now,
		}).
		OnDuplicateKeyUpdate().
		ValueOnUpdateAge(sql.NullInt64{
			Valid: true,
			Int64: 17,
		})
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expr := regexp.MustCompile("^INSERT INTO `user` \\(.*\\) VALUES\\(\\?,\\?,\\?,\\?\\) ")
	gotSuffix := expr.ReplaceAllString(query, "")
	expectedSuffix1 := "ON DUPLICATE KEY UPDATE `age` = ?, `updated_at` = VALUES(`updated_at`);"
	expectedSuffix2 := "ON DUPLICATE KEY UPDATE `updated_at` = VALUES(`updated_at`), `age` = ?;"
	if gotSuffix != expectedSuffix1 && gotSuffix != expectedSuffix2 {
		t.Error("unexpected suffix:", gotSuffix)
	}
	if len(args) != 5 {
		t.Error("args is too many:", len(args))
	}
}

func TestBulkInsert(t *testing.T) {
	items := example.NewUserItemSQL().BulkInsert()
	for i := 1; i <= 10; i++ {
		q := example.NewUserItemSQL().Insert().
			ValueUserID(42).
			ValueItemID(strconv.Itoa(i))
		items.Append(q)
	}
	query, vs, err := items.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	expected := "INSERT INTO `user_item` (`item_id`,`user_id`) VALUES (?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?),(?,?);"
	if query != expected {
		t.Error("query is not match:", query)
	}
	if !reflect.DeepEqual(vs, []interface{}{"1", uint64(42), "2", uint64(42), "3", uint64(42), "4", uint64(42), "5", uint64(42), "6", uint64(42), "7", uint64(42), "8", uint64(42), "9", uint64(42), "10", uint64(42)}) {
		t.Errorf("vs is not valid: %+v", vs)
	}
}

func TestBulkInsertWithOnDuplicateKeyUpdate(t *testing.T) {
	items := example.NewUserItemSQL().BulkInsert()
	items.Append(
		example.NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("1").ValueIsUsed(true),
		example.NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("2").ValueIsUsed(true),
	)

	now := sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	query, vs, err := items.
		OnDuplicateKeyUpdate().
		SameOnUpdateIsUsed().
		ValueOnUpdateUsedAt(now).
		ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	bulkInsertQuery := "INSERT INTO `user_item` (`is_used`,`item_id`,`user_id`) VALUES (?,?,?),(?,?,?) "
	expected1 := bulkInsertQuery + "ON DUPLICATE KEY UPDATE `is_used` = VALUES(`is_used`), `used_at` = ?;"
	expected2 := bulkInsertQuery + "ON DUPLICATE KEY UPDATE `used_at` = ?, `is_used` = VALUES(`is_used`);"
	if query != expected1 && query != expected2 {
		t.Error("query is not match:", query)
	}
	if !reflect.DeepEqual(vs, []interface{}{true, "1", uint64(42), true, "2", uint64(42), now}) {
		t.Errorf("vs is not valid: %+v", vs)
	}
}

func TestDelete(t *testing.T) {
	q := example.NewUserSQL().Delete().Name("hogehoge")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "DELETE FROM `user` WHERE `name` = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge"}) {
		t.Error("unexpected args:", args)
	}
}

func TestDelete__In(t *testing.T) {
	q := example.NewUserSQL().Delete().NameIn("hogehoge", "fugafuga")
	query, args, err := q.ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "DELETE FROM `user` WHERE `name` IN(?,?);" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"hogehoge", "fugafuga"}) {
		t.Error("unexpected args:", args)
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

func TestORM__Plugin__Count(t *testing.T) {
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
