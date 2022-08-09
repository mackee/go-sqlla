//go:build withmysql

package example

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mackee/go-sqlla/v2"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB

//go:generate genddl -outpath=./mysql.sql -driver=mysql

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{
		"MYSQL_ROOT_PASSWORD=secret",
		"MYSQL_DATABASE=test",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/test?parseTime=true", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	schemaFile, err := os.Open("./mysql.sql")
	if err != nil {
		log.Fatal("cannot open schema file error:", err)
	}

	b, err := ioutil.ReadAll(schemaFile)
	if err != nil {
		log.Fatal("cannot read schema file error:", err)
	}

	stmts := strings.Split(string(b), ";")
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal("cannot load schema error:", err)
		}
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestInsertWithoutSelect__WithMySQL(t *testing.T) {
	ctx := context.Background()
	now := time.Now()

	q1 := NewUserExternalSQL().Insert().
		ValueID(42).
		ValueUserID(42).
		ValueCreatedAt(now).
		ValueUpdatedAt(now)
	_, err := q1.ExecContextWithoutSelect(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func TestInsertOnDuplicateKeyUpdate__WithMySQL(t *testing.T) {
	ctx := context.Background()
	now1 := time.Now()

	q1 := NewUserSQL().Insert().
		ValueName("hogehoge").
		ValueRate(3.14).
		ValueIconImage([]byte{}).
		ValueAge(sql.NullInt64{Valid: true, Int64: 17}).
		ValueUpdatedAt(mysql.NullTime{Valid: true, Time: now1})
	query, args, _ := q1.ToSql()
	t.Logf("query=%s, args=%+v", query, args)
	r1, err := q1.ExecContext(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	now2 := now1.Add(1 * time.Second)

	q2 := NewUserSQL().Insert().
		ValueName("hogehoge").
		ValueAge(sql.NullInt64{Valid: true, Int64: 17}).
		ValueIconImage([]byte{}).
		ValueUpdatedAt(mysql.NullTime{Valid: true, Time: now2}).
		OnDuplicateKeyUpdate().
		RawValueOnUpdateAge(sqlla.SetMapRawValue("`age` + 1"))
	r2, err := q2.ExecContext(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if r2.Rate != 3.14 {
		t.Fatal("rate is not match:", r2.Rate)
	}
	if r2.Age.Int64 != 18 {
		t.Fatal("age does not incremented:", r2.Age.Int64)
	}
	if r2.UpdatedAt.Time.Unix() <= r1.UpdatedAt.Time.Unix() {
		t.Fatal("updated_at does not updated:", r1.UpdatedAt.Time.Unix(), r2.UpdatedAt.Time.Unix())
	}
}

func TestBulkInsert__WithMySQL(t *testing.T) {
	ctx := context.Background()

	if _, err := NewUserItemSQL().Delete().ExecContext(ctx, db); err != nil {
		t.Fatal("unexpected error:", err)
	}

	items := NewUserItemSQL().BulkInsert()
	items.Append(
		NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("1").ValueIsUsed(true),
		NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("2").ValueIsUsed(true),
	)

	if _, err := items.ExecContext(ctx, db); err != nil {
		t.Fatal("unexpected error:", err)
	}

	uis, err := NewUserItemSQL().Select().AllContext(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	for i, ui := range uis {
		if ui.UserId != 42 {
			t.Error("UserId is not match:", ui.UserId)
		}
		if ui.ItemId != strconv.Itoa(i+1) {
			t.Errorf("ItemId is not match: index=%d, got=%s", i, ui.ItemId)
		}
		if !ui.IsUsed {
			t.Error("IsUsed is false")
		}
	}
}

func TestBulkInsertOnDuplicateKeyUpdate__WithMySQL(t *testing.T) {
	ctx := context.Background()

	if _, err := NewUserItemSQL().Delete().ExecContext(ctx, db); err != nil {
		t.Fatal("unexpected error:", err)
	}

	items := NewUserItemSQL().BulkInsert()
	items.Append(
		NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("1").ValueIsUsed(false),
		NewUserItemSQL().Insert().ValueUserID(42).ValueItemID("2").ValueIsUsed(false),
	)

	if _, err := items.ExecContext(ctx, db); err != nil {
		t.Fatal("unexpected error:", err)
	}

	uis, err := NewUserItemSQL().Select().AllContext(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	uitems := NewUserItemSQL().BulkInsert()
	for _, ui := range uis {
		uitems.Append(
			NewUserItemSQL().Insert().
				ValueID(ui.Id).
				ValueUserID(42).
				ValueItemID(ui.ItemId).
				ValueIsUsed(true),
		)
	}
	uitems.Append(
		NewUserItemSQL().Insert().
			ValueID(uis[len(uis)-1].Id + 1).
			ValueUserID(42).
			ValueItemID("3").
			ValueIsUsed(true),
	)
	now := time.Now()
	dup := uitems.OnDuplicateKeyUpdate().
		SameOnUpdateIsUsed().
		ValueOnUpdateUsedAt(sql.NullTime{
			Valid: true,
			Time:  now,
		})

	if _, err := dup.ExecContext(ctx, db); err != nil {
		t.Fatal("unexpected error:", err)
	}

	uuis, err := NewUserItemSQL().Select().OrderByID(sqlla.Asc).AllContext(ctx, db)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	for i, ui := range uuis {
		if !ui.IsUsed {
			t.Errorf("IsUsed is false: index=%d", i)
		}
		switch i {
		case 0, 1:
			if !ui.UsedAt.Valid {
				t.Errorf("UsedAt is not valid: index=%d", i)
			}
		case 2:
			if ui.UsedAt.Valid {
				t.Errorf("UsedAt is valid: index=%d", i)
			}
		}
	}

}
