//go:build withpostgresql

package postgresql_test

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mackee/go-sqlla/_example/postgresql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pgvector/pgvector-go"
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "pgvector/pgvector",
		Tag:        "pg17",
		Env: []string{
			"POSTGRES_USER=sqlla_test",
			"POSTGRES_PASSWORD=secret",
		},
	},
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		port, err := strconv.Atoi(resource.GetPort("5432/tcp"))
		if err != nil {
			return fmt.Errorf("cannot convert port to int: %w", err)
		}
		db, err = sql.Open("pgx", fmt.Sprintf("host=localhost port=%d user=sqlla_test password=secret dbname=sqlla_test sslmode=disable", port))
		if err != nil {
			return fmt.Errorf("cannot open database: %w", err)
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS vector;"); err != nil {
		log.Fatalf("Could not create extension: %s", err)
	}

	schemaFile, err := os.Open("./postgresql.sql")
	if err != nil {
		log.Fatal("cannot open schema file error:", err)
	}

	b, err := io.ReadAll(schemaFile)
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

func cleanupDB(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "TRUNCATE accounts RESTART IDENTITY"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "TRUNCATE identities RESTART IDENTITY"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "TRUNCATE groups RESTART IDENTITY"); err != nil {
		t.Fatal(err)
	}
}

func TestDB(t *testing.T) {
	postgresql.FixedNow = sampleDate

	tcs := testCases()
	opts := cmp.Options{
		cmpopts.EquateEmpty(),
		cmpopts.IgnoreUnexported(pgvector.Vector{}),
	}
	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			defer cleanupDB(t)
			tc.assert(t, opts...)
		})
	}
}
