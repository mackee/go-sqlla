package postgresql_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mackee/go-sqlla/_example/postgresql"
	"github.com/mackee/go-sqlla/v2"
	"github.com/pgvector/pgvector-go"
)

var db *sql.DB

type toSqler interface {
	ToSql() (string, []any, error)
}

type testCase interface {
	Name() string
	assert(t *testing.T, opts ...cmp.Option)
}

type testCaseWithToQueryTestCase interface {
	testCase
	toQueryTestCase() queryTestCase
}

type queryTestCase struct {
	name     string
	query    toSqler
	expected string
	vs       []any
}

func (q queryTestCase) Name() string {
	return q.name
}

func (q queryTestCase) assert(t *testing.T, opts ...cmp.Option) {
	t.Helper()

	sql, vs, err := q.query.ToSql()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sql != q.expected {
		t.Errorf("expected: \n%s\n, but got: \n%s", q.expected, sql)
	}
	if len(vs) != len(q.vs) {
		t.Errorf("expected: %v, but got: %v", q.vs, vs)
	}
	if diff := cmp.Diff(vs, q.vs, opts...); diff != "" {
		t.Errorf("has diff: %s", diff)
	}
}

type selectQuery[T any] interface {
	toSqler
	SingleContext(ctx context.Context, db sqlla.DB) (T, error)
	AllContext(ctx context.Context, db sqlla.DB) ([]T, error)
}

type selectQueryTestCase[T any] struct {
	name           string
	query          selectQuery[T]
	expected       string
	vs             []any
	expectedResult []T
}

func (s selectQueryTestCase[T]) Name() string {
	return s.name
}

func (s selectQueryTestCase[T]) toQueryTestCase() queryTestCase {
	return queryTestCase{
		name:     s.name,
		query:    s.query,
		expected: s.expected,
		vs:       s.vs,
	}
}

func (s selectQueryTestCase[T]) assert(t *testing.T, opts ...cmp.Option) {
	t.Helper()

	ctx := context.Background()
	switch len(s.expectedResult) {
	case 0:
		if _, err := s.query.SingleContext(ctx, db); errors.Is(err, sql.ErrNoRows) {
			return
		} else {
			t.Fatalf("unexpected error: %v", err)
		}
	case 1:
		got, err := s.query.SingleContext(ctx, db)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := cmp.Diff(got, s.expectedResult[0], opts...); diff != "" {
			t.Errorf("has diff: %s", diff)
		}
	default:
		got, err := s.query.AllContext(ctx, db)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if diff := cmp.Diff(got, s.expectedResult, opts...); diff != "" {
			t.Errorf("has diff: %s", diff)
		}
	}
}

type insertQuery[T any] interface {
	toSqler
	ExecContext(ctx context.Context, db sqlla.DB) (T, error)
}

type insertQueryTestCase[T any] struct {
	name           string
	query          insertQuery[T]
	expected       string
	vs             []any
	expectedResult T
}

func (i insertQueryTestCase[T]) Name() string {
	return i.name
}

func (i insertQueryTestCase[T]) toQueryTestCase() queryTestCase {
	return queryTestCase{
		name:     i.name,
		query:    i.query,
		expected: i.expected,
		vs:       i.vs,
	}
}

func (i insertQueryTestCase[T]) assert(t *testing.T, opts ...cmp.Option) {
	t.Helper()

	ctx := context.Background()
	got, err := i.query.ExecContext(ctx, db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(got, i.expectedResult, opts...); diff != "" {
		t.Errorf("has diff: %s", diff)
	}
}

type bulkInsertQuery[T any] interface {
	toSqler
	ExecContext(ctx context.Context, db sqlla.DB) ([]T, error)
}

type bulkInsertQueryTestCase[T any] struct {
	name           string
	query          bulkInsertQuery[T]
	expected       string
	vs             []any
	expectedResult []T
}

func (b bulkInsertQueryTestCase[T]) Name() string {
	return b.name
}

func (b bulkInsertQueryTestCase[T]) toQueryTestCase() queryTestCase {
	return queryTestCase{
		name:     b.name,
		query:    b.query,
		expected: b.expected,
		vs:       b.vs,
	}
}

func (b bulkInsertQueryTestCase[T]) assert(t *testing.T, opts ...cmp.Option) {
	t.Helper()

	ctx := context.Background()
	got, err := b.query.ExecContext(ctx, db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(got, b.expectedResult, opts...); diff != "" {
		t.Errorf("has diff: %s", diff)
	}
}

type updateQuery[T any] interface {
	toSqler
	ExecContext(ctx context.Context, db sqlla.DB) ([]T, error)
}

type updateQueryTestCase[T any] struct {
	name           string
	query          updateQuery[T]
	expected       string
	vs             []any
	setup          func(t *testing.T, db sqlla.DB)
	expectedResult []T
}

func (u updateQueryTestCase[T]) Name() string {
	return u.name
}

func (u updateQueryTestCase[T]) toQueryTestCase() queryTestCase {
	return queryTestCase{
		name:     u.name,
		query:    u.query,
		expected: u.expected,
		vs:       u.vs,
	}
}

func (u updateQueryTestCase[T]) assert(t *testing.T, opts ...cmp.Option) {
	t.Helper()

	ctx := context.Background()
	if u.setup != nil {
		u.setup(t, db)
	}
	got, err := u.query.ExecContext(ctx, db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(got, u.expectedResult, opts...); diff != "" {
		t.Errorf("has diff: %s", diff)
	}
}

type deleteQuery interface {
	toSqler
	ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error)
}

type deleteQueryTestCase struct {
	name     string
	query    deleteQuery
	expected string
	vs       []any
}

func (d deleteQueryTestCase) Name() string {
	return d.name
}

func (d deleteQueryTestCase) toQueryTestCase() queryTestCase {
	return queryTestCase{
		name:     d.name,
		query:    d.query,
		expected: d.expected,
		vs:       d.vs,
	}
}

func (d deleteQueryTestCase) assert(t *testing.T, _ ...cmp.Option) {
	t.Helper()

	ctx := context.Background()
	if _, err := d.query.ExecContext(ctx, db); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

var (
	sampleDate   = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	sampleVector = pgvector.NewVector([]float32{1, 2, 3})
)

func testCases() []testCaseWithToQueryTestCase {
	postgresql.FixedNow = sampleDate

	return []testCaseWithToQueryTestCase{
		selectQueryTestCase[postgresql.Account]{
			name:     "select all",
			query:    postgresql.NewAccountSQL().Select(),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts";`,
			vs:       nil,
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by id",
			query:    postgresql.NewAccountSQL().Select().ID(42),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" = $1;`,
			vs:       []any{int64(42)},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by range",
			query:    postgresql.NewAccountSQL().Select().CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "created_at" < $1;`,
			vs:       []any{sampleDate},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by complex where",
			query:    postgresql.NewAccountSQL().Select().ID(42).CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" = $1 AND "created_at" < $2;`,
			vs:       []any{int64(42), sampleDate},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by IN(...)",
			query:    postgresql.NewAccountSQL().Select().IDIn(6, 28, 496, 8128),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" IN($1,$2,$3,$4);`,
			vs:       []any{int64(6), int64(28), int64(496), int64(8128)},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by IN(...) and range",
			query:    postgresql.NewAccountSQL().Select().IDIn(6, 28, 496, 8128).CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" IN($1,$2,$3,$4) AND "created_at" < $5;`,
			vs:       []any{int64(6), int64(28), int64(496), int64(8128), sampleDate},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select by range and IN(...)",
			query:    postgresql.NewAccountSQL().Select().CreatedAt(sampleDate, sqlla.OpLess).IDIn(6, 28, 496, 8128),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "created_at" < $1 AND "id" IN($2,$3,$4,$5);`,
			vs:       []any{sampleDate, int64(6), int64(28), int64(496), int64(8128)},
		},
		selectQueryTestCase[postgresql.Account]{
			name: "select by additional where clause",
			query: postgresql.NewAccountSQL().Select().AdditionalWhereClause(func(offset int) (string, int, []any) {
				return fmt.Sprintf(`WHERE "name" = $%d`, offset+1), offset + 1, []any{"foo"}
			}),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "name" = $1;`,
			vs:       []any{"foo"},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select with alias",
			query:    postgresql.NewAccountSQL().Select().TableAlias("a").ID(42),
			expected: `SELECT "a"."id", "a"."name", "a"."embedding", "a"."created_at", "a"."updated_at" FROM "accounts" AS "a" WHERE "a"."id" = $1;`,
			vs:       []any{int64(42)},
		},
		selectQueryTestCase[postgresql.Account]{
			name: "select with alias and join clause",
			query: postgresql.NewAccountSQL().Select().TableAlias("a").
				JoinClause("INNER JOIN identities i ON i.account_id = a.id").
				ID(42).
				AdditionalWhereClause(func(offset int) (string, int, []any) {
					return fmt.Sprintf("AND i.email = $%d", offset+1), offset + 1, []any{"hogehoge@example.com"}
				}),
			expected: `SELECT "a"."id", "a"."name", "a"."embedding", "a"."created_at", "a"."updated_at" FROM "accounts" AS "a" INNER JOIN identities i ON i.account_id = a.id WHERE "a"."id" = $1 AND i.email = $2;`,
			vs:       []any{int64(42), "hogehoge@example.com"},
		},
		selectQueryTestCase[postgresql.Account]{
			name: "select by or expression",
			query: postgresql.NewAccountSQL().Select().Or(
				postgresql.NewAccountSQL().Select().Name("foo"),
				postgresql.NewAccountSQL().Select().Name("bar"),
			),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE (( "name" = $1 ) OR ( "name" = $2 ));`,
			vs:       []any{"foo", "bar"},
		},
		selectQueryTestCase[postgresql.Account]{
			name: "select by cosine simularity",
			query: postgresql.NewAccountSQL().Select().
				Embedding(
					sampleVector,
					sqlla.OpPgvectorCosine,
					sqlla.NewOperatorAndValue(sqlla.OpGreater, 0.8),
				).
				ID(42),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE ( "embedding" <=> $1 ) > $2 AND "id" = $3;`,
			vs:       []any{sampleVector, float64(0.8), int64(42)},
		},
		selectQueryTestCase[postgresql.Group]{
			name: "select with null column",
			query: postgresql.NewGroupSQL().Select().
				SubLeaderAccountID(42),
			expected: `SELECT "id", "name", "leader_account_id", "sub_leader_account_id", "child_group_id", "created_at", "updated_at" FROM "groups" WHERE "sub_leader_account_id" = $1;`,
			vs:       []any{int64(42)},
		},
		selectQueryTestCase[postgresql.Group]{
			name: "select with null column is null",
			query: postgresql.NewGroupSQL().Select().
				ID(42).
				SubLeaderAccountIDIsNull().
				ChildGroupID(28),
			expected: `SELECT "id", "name", "leader_account_id", "sub_leader_account_id", "child_group_id", "created_at", "updated_at" FROM "groups" WHERE "id" = $1 AND "sub_leader_account_id" IS NULL AND "child_group_id" = $2;`,
			vs:       []any{int64(42), int64(28)},
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select with limit",
			query:    postgresql.NewAccountSQL().Select().Limit(10),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" LIMIT 10;`,
			vs:       nil,
		},
		selectQueryTestCase[postgresql.Account]{
			name:     "select order by",
			query:    postgresql.NewAccountSQL().Select().OrderByID(sqlla.Asc),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" ORDER BY "id" ASC;`,
			vs:       nil,
		},
		selectQueryTestCase[postgresql.Account]{
			name: "select order by consine simularity",
			query: postgresql.NewAccountSQL().Select().
				ID(42).
				OrderByEmbedding(
					sqlla.NewOrderWithOperator(
						sqlla.OpPgvectorCosine,
						sampleVector,
						sqlla.Desc,
					),
				).
				Limit(10),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" = $1 ORDER BY "embedding" <=> $2 DESC LIMIT 10;`,
			vs:       []any{int64(42), sampleVector},
		},
		insertQueryTestCase[postgresql.Account]{
			name: "insert",
			query: postgresql.NewAccountSQL().Insert().
				ValueName("foo").
				ValueEmbedding(sampleVector),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs:       []any{sampleDate, sampleVector, "foo", sampleDate},
			expectedResult: postgresql.Account{
				ID:        1,
				Name:      "foo",
				Embedding: sampleVector,
				CreatedAt: sampleDate,
				UpdatedAt: sampleDate,
			},
		},
		insertQueryTestCase[postgresql.Group]{
			name: "insert with null column",
			query: postgresql.NewGroupSQL().Insert().
				ValueName("foo").
				ValueLeaderAccountID(42).
				ValueSubLeaderAccountID(28).
				ValueChildGroupIDIsNull().
				ValueCreatedAt(sampleDate).
				ValueUpdatedAt(sampleDate),
			expected: `INSERT INTO "groups" ("child_group_id","created_at","leader_account_id","name","sub_leader_account_id","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id", "name", "leader_account_id", "sub_leader_account_id", "child_group_id", "created_at", "updated_at";`,
			vs:       []any{sql.Null[int64]{}, sampleDate, int64(42), "foo", int64(28), sampleDate},
			expectedResult: postgresql.Group{
				ID:                 1,
				Name:               "foo",
				LeaderAccountID:    42,
				SubLeaderAccountID: sql.Null[postgresql.AccountID]{Valid: true, V: 28},
				ChildGroupID:       sql.Null[postgresql.GroupID]{},
				CreatedAt:          sampleDate,
				UpdatedAt:          sampleDate,
			},
		},
		insertQueryTestCase[postgresql.Account]{
			name: "insert on conflict do nothing",
			query: postgresql.NewAccountSQL().Insert().
				ValueName("foo").
				ValueEmbedding(sampleVector).
				OnConflictDoNothing(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4) ON CONFLICT DO NOTHING RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs:       []any{sampleDate, sampleVector, "foo", sampleDate},
			expectedResult: postgresql.Account{
				ID:        1,
				Name:      "foo",
				Embedding: sampleVector,
				CreatedAt: sampleDate,
				UpdatedAt: sampleDate,
			},
		},
		insertQueryTestCase[postgresql.Account]{
			name: "insert on conflict do update",
			query: postgresql.NewAccountSQL().Insert().
				ValueName("foo").
				ValueEmbedding(sampleVector).
				OnConflictDoUpdate("id").
				ValueOnUpdateName("powawa").
				SameOnUpdateEmbedding(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4) ON CONFLICT (id) DO UPDATE SET "embedding" = "excluded"."embedding", "name" = $5 RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs:       []any{sampleDate, sampleVector, "foo", sampleDate, "powawa"},
			expectedResult: postgresql.Account{
				ID:        1,
				Name:      "foo",
				Embedding: sampleVector,
				CreatedAt: sampleDate,
				UpdatedAt: sampleDate,
			},
		},
		bulkInsertQueryTestCase[postgresql.Account]{
			name: "bulk insert",
			query: func() bulkInsertQuery[postgresql.Account] {
				bi := postgresql.NewAccountSQL().BulkInsert()
				for _, s := range []string{"foo", "bar", "baz"} {
					q := postgresql.NewAccountSQL().Insert().
						ValueName(s).
						ValueEmbedding(sampleVector)
					bi.Append(q)
				}
				return bi
			}(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12) RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs: []any{
				sampleDate,
				sampleVector,
				"foo",
				sampleDate,
				sampleDate,
				sampleVector,
				"bar",
				sampleDate,
				sampleDate,
				sampleVector,
				"baz",
				sampleDate,
			},
			expectedResult: []postgresql.Account{
				{ID: 1, Name: "foo", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 2, Name: "bar", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 3, Name: "baz", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
			},
		},
		bulkInsertQueryTestCase[postgresql.Account]{
			name: "bulk insert with on conflict do nothing",
			query: func() bulkInsertQuery[postgresql.Account] {
				bi := postgresql.NewAccountSQL().BulkInsert()
				for _, s := range []string{"foo", "bar", "baz"} {
					q := postgresql.NewAccountSQL().Insert().
						ValueName(s).
						ValueEmbedding(sampleVector)
					bi.Append(q)
				}
				return bi.OnConflictDoNothing()
			}(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12) ON CONFLICT DO NOTHING RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs: []any{
				sampleDate,
				sampleVector,
				"foo",
				sampleDate,
				sampleDate,
				sampleVector,
				"bar",
				sampleDate,
				sampleDate,
				sampleVector,
				"baz",
				sampleDate,
			},
			expectedResult: []postgresql.Account{
				{ID: 1, Name: "foo", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 2, Name: "bar", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 3, Name: "baz", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
			},
		},
		bulkInsertQueryTestCase[postgresql.Account]{
			name: "bulk insert with on conflict do update",
			query: func() bulkInsertQuery[postgresql.Account] {
				bi := postgresql.NewAccountSQL().BulkInsert()
				for id, s := range []string{"foo", "bar", "baz"} {
					q := postgresql.NewAccountSQL().Insert().
						ValueID(postgresql.AccountID(id + 1)).
						ValueName(s).
						ValueEmbedding(sampleVector)
					bi.Append(q)
				}
				return bi.OnConflictDoUpdate("id").
					ValueOnUpdateName("powawa").
					SameOnUpdateEmbedding()
			}(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","id","name","updated_at") VALUES ($1,$2,$3,$4,$5),($6,$7,$8,$9,$10),($11,$12,$13,$14,$15) ON CONFLICT (id) DO UPDATE SET "embedding" = "excluded"."embedding", "name" = $16 RETURNING "id", "name", "embedding", "created_at", "updated_at";`,
			vs: []any{
				sampleDate,
				sampleVector,
				int64(1),
				"foo",
				sampleDate,
				sampleDate,
				sampleVector,
				int64(2),
				"bar",
				sampleDate,
				sampleDate,
				sampleVector,
				int64(3),
				"baz",
				sampleDate,
				"powawa",
			},
			expectedResult: []postgresql.Account{
				{ID: 1, Name: "foo", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 2, Name: "bar", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
				{ID: 3, Name: "baz", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
			},
		},
		updateQueryTestCase[postgresql.Account]{
			name: "update",
			setup: func(t *testing.T, db sqlla.DB) {
				if _, err := postgresql.NewAccountSQL().Insert().
					ValueID(42).
					ValueName("foo").
					ValueEmbedding(sampleVector).
					ValueCreatedAt(sampleDate).
					ValueUpdatedAt(sampleDate).
					ExecContextWithoutSelect(t.Context(), db); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
			query: postgresql.NewAccountSQL().Update().
				SetName("bar").
				SetEmbedding(sampleVector).
				WhereID(42),
			expected: `UPDATE "accounts" SET "embedding" = $1, "name" = $2, "updated_at" = $3 WHERE "id" = $4;`,
			vs:       []any{sampleVector, "bar", sampleDate, int64(42)},
			expectedResult: []postgresql.Account{
				{ID: 42, Name: "bar", Embedding: sampleVector, CreatedAt: sampleDate, UpdatedAt: sampleDate},
			},
		},
		updateQueryTestCase[postgresql.Group]{
			name: "update with set null",
			setup: func(t *testing.T, db sqlla.DB) {
				if _, err := postgresql.NewGroupSQL().Insert().
					ValueID(111).
					ValueName("foo").
					ValueLeaderAccountID(42).
					ValueSubLeaderAccountID(28).
					ValueChildGroupID(43).
					ValueCreatedAt(sampleDate).
					ValueUpdatedAt(sampleDate).
					ExecContextWithoutSelect(t.Context(), db); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
			query: postgresql.NewGroupSQL().Update().
				SetLeaderAccountID(42).
				SetSubLeaderAccountID(28).
				SetChildGroupIDToNull().
				WhereID(111),
			expected: `UPDATE "groups" SET "child_group_id" = $1, "leader_account_id" = $2, "sub_leader_account_id" = $3 WHERE "id" = $4;`,
			vs:       []any{sql.Null[int64]{}, int64(42), int64(28), int64(111)},
			expectedResult: []postgresql.Group{
				{ID: 111, Name: "foo", LeaderAccountID: 42, SubLeaderAccountID: sql.Null[postgresql.AccountID]{Valid: true, V: 28}, ChildGroupID: sql.Null[postgresql.GroupID]{}, CreatedAt: sampleDate, UpdatedAt: sampleDate},
			},
		},
		deleteQueryTestCase{
			name:     "delete with where",
			query:    postgresql.NewAccountSQL().Delete().ID(42),
			expected: `DELETE FROM "accounts" WHERE "id" = $1;`,
			vs:       []any{int64(42)},
		},
		deleteQueryTestCase{
			name: "delete by cosine simularity",
			query: postgresql.NewAccountSQL().Delete().
				Embedding(sampleVector, sqlla.OpPgvectorCosine, sqlla.NewOperatorAndValue(sqlla.OpGreater, 0.8)),
			expected: `DELETE FROM "accounts" WHERE ( "embedding" <=> $1 ) > $2;`,
			vs:       []any{sampleVector, float64(0.8)},
		},
		deleteQueryTestCase{
			name: "delete with null column",
			query: postgresql.NewGroupSQL().Delete().
				SubLeaderAccountID(42).
				ChildGroupIDIsNull().
				ID(111),
			expected: `DELETE FROM "groups" WHERE "sub_leader_account_id" = $1 AND "child_group_id" IS NULL AND "id" = $2;`,
			vs:       []any{int64(42), int64(111)},
		},
	}
}

func TestAccountToSql(t *testing.T) {
	postgresql.FixedNow = sampleDate

	tcs := testCases()
	opts := cmp.Options{
		cmpopts.IgnoreUnexported(pgvector.Vector{}),
	}
	for _, tc := range tcs {
		t.Run(tc.Name(), func(t *testing.T) {
			tc.toQueryTestCase().assert(t, opts...)
		})
	}
}
