package postgresql_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/mackee/go-sqlla/_example/postgresql"
	"github.com/mackee/go-sqlla/v2"
	"github.com/pgvector/pgvector-go"
)

type toSqler interface {
	ToSql() (string, []any, error)
}

var (
	sampleDate   = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	sampleVector = pgvector.NewVector([]float32{1, 2, 3})
)

func TestAccountToSql(t *testing.T) {
	postgresql.FixedNow = sampleDate
	tcs := []struct {
		name     string
		query    toSqler
		expected string
		vs       []any
	}{
		{
			name:     "select all",
			query:    postgresql.NewAccountSQL().Select(),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts";`,
			vs:       nil,
		},
		{
			name:     "select by id",
			query:    postgresql.NewAccountSQL().Select().ID(42),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" = $1;`,
			vs:       []any{int64(42)},
		},
		{
			name:     "select by range",
			query:    postgresql.NewAccountSQL().Select().CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "created_at" < $1;`,
			vs:       []any{sampleDate},
		},
		{
			name:     "select by complex where",
			query:    postgresql.NewAccountSQL().Select().ID(42).CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" = $1 AND "created_at" < $2;`,
			vs:       []any{int64(42), sampleDate},
		},
		{
			name:     "select by IN(...)",
			query:    postgresql.NewAccountSQL().Select().IDIn(6, 28, 496, 8128),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" IN($1,$2,$3,$4);`,
			vs:       []any{int64(6), int64(28), int64(496), int64(8128)},
		},
		{
			name:     "select by IN(...) and range",
			query:    postgresql.NewAccountSQL().Select().IDIn(6, 28, 496, 8128).CreatedAt(sampleDate, sqlla.OpLess),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "id" IN($1,$2,$3,$4) AND "created_at" < $5;`,
			vs:       []any{int64(6), int64(28), int64(496), int64(8128), sampleDate},
		},
		{
			name:     "select by range and IN(...)",
			query:    postgresql.NewAccountSQL().Select().CreatedAt(sampleDate, sqlla.OpLess).IDIn(6, 28, 496, 8128),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "created_at" < $1 AND "id" IN($2,$3,$4,$5);`,
			vs:       []any{sampleDate, int64(6), int64(28), int64(496), int64(8128)},
		},
		{
			name: "select by additional where clause",
			query: postgresql.NewAccountSQL().Select().AdditionalWhereClause(func(offset int) (string, int, []any) {
				return fmt.Sprintf(`WHERE "name" = $%d`, offset+1), offset + 1, []any{"foo"}
			}),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE "name" = $1;`,
			vs:       []any{"foo"},
		},
		{
			name:     "select with alias",
			query:    postgresql.NewAccountSQL().Select().TableAlias("a").ID(42),
			expected: `SELECT "a"."id", "a"."name", "a"."embedding", "a"."created_at", "a"."updated_at" FROM "accounts" AS "a" WHERE "a"."id" = $1;`,
			vs:       []any{int64(42)},
		},
		{
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
		{
			name: "select by or expression",
			query: postgresql.NewAccountSQL().Select().Or(
				postgresql.NewAccountSQL().Select().Name("foo"),
				postgresql.NewAccountSQL().Select().Name("bar"),
			),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" WHERE (( "name" = $1 ) OR ( "name" = $2 ));`,
			vs:       []any{"foo", "bar"},
		},
		{
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
		{
			name: "select with null column",
			query: postgresql.NewGroupSQL().Select().
				SubLeaderAccountID(42),
			expected: `SELECT "id", "name", "leader_account_id", "sub_leader_account_id", "child_group_id", "created_at", "updated_at" FROM "groups" WHERE "sub_leader_account_id" = $1;`,
			vs:       []any{int64(42)},
		},
		{
			name: "select with null column is null",
			query: postgresql.NewGroupSQL().Select().
				ID(42).
				SubLeaderAccountIDIsNull().
				ChildGroupID(28),
			expected: `SELECT "id", "name", "leader_account_id", "sub_leader_account_id", "child_group_id", "created_at", "updated_at" FROM "groups" WHERE "id" = $1 AND "sub_leader_account_id" IS NULL AND "child_group_id" = $2;`,
			vs:       []any{int64(42), int64(28)},
		},
		{
			name:     "select with limit",
			query:    postgresql.NewAccountSQL().Select().Limit(10),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" LIMIT 10;`,
			vs:       nil,
		},
		{
			name:     "select order by",
			query:    postgresql.NewAccountSQL().Select().OrderByID(sqlla.Asc),
			expected: `SELECT "id", "name", "embedding", "created_at", "updated_at" FROM "accounts" ORDER BY "id" ASC;`,
			vs:       nil,
		},
		{
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
		{
			name: "insert",
			query: postgresql.NewAccountSQL().Insert().
				ValueName("foo").
				ValueEmbedding(sampleVector),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id";`,
			vs:       []any{sampleDate, sampleVector, "foo", sampleDate},
		},
		{
			name: "insert with null column",
			query: postgresql.NewGroupSQL().Insert().
				ValueName("foo").
				ValueLeaderAccountID(42).
				ValueSubLeaderAccountID(28).
				ValueChildGroupIDIsNull().
				ValueCreatedAt(sampleDate).
				ValueUpdatedAt(sampleDate),
			expected: `INSERT INTO "groups" ("child_group_id","created_at","leader_account_id","name","sub_leader_account_id","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id";`,
			vs:       []any{sql.Null[int64]{}, sampleDate, int64(42), "foo", int64(28), sampleDate},
		},
		{
			name: "bulk insert",
			query: func() toSqler {
				bi := postgresql.NewAccountSQL().BulkInsert()
				for _, s := range []string{"foo", "bar", "baz"} {
					q := postgresql.NewAccountSQL().Insert().
						ValueName(s).
						ValueEmbedding(sampleVector)
					bi.Append(q)
				}
				return bi
			}(),
			expected: `INSERT INTO "accounts" ("created_at","embedding","name","updated_at") VALUES ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12) RETURNING "id";`,
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
		},
		{
			name: "update",
			query: postgresql.NewAccountSQL().Update().
				SetName("bar").
				SetEmbedding(sampleVector).
				WhereID(42),
			expected: `UPDATE "accounts" SET "embedding" = $1, "name" = $2, "updated_at" = $3 WHERE "id" = $4;`,
			vs:       []any{sampleVector, "bar", sampleDate, int64(42)},
		},
		{
			name: "update with set null",
			query: postgresql.NewGroupSQL().Update().
				SetLeaderAccountID(42).
				SetSubLeaderAccountID(28).
				SetChildGroupIDToNull().
				WhereID(111),
			expected: `UPDATE "groups" SET "child_group_id" = $1, "leader_account_id" = $2, "sub_leader_account_id" = $3 WHERE "id" = $4;`,
			vs:       []any{sql.Null[int64]{}, int64(42), int64(28), int64(111)},
		},
		{
			name:     "delete with where",
			query:    postgresql.NewAccountSQL().Delete().ID(42),
			expected: `DELETE FROM "accounts" WHERE "id" = $1;`,
			vs:       []any{int64(42)},
		},
		{
			name: "delete by cosine simularity",
			query: postgresql.NewAccountSQL().Delete().
				Embedding(sampleVector, sqlla.OpPgvectorCosine, sqlla.NewOperatorAndValue(sqlla.OpGreater, 0.8)),
			expected: `DELETE FROM "accounts" WHERE ( "embedding" <=> $1 ) > $2;`,
			vs:       []any{sampleVector, float64(0.8)},
		},
		{
			name: "delete with null column",
			query: postgresql.NewGroupSQL().Delete().
				SubLeaderAccountID(42).
				ChildGroupIDIsNull().
				ID(111),
			expected: `DELETE FROM "groups" WHERE "sub_leader_account_id" = $1 AND "child_group_id" IS NULL AND "id" = $2;`,
			vs:       []any{int64(42), int64(111)},
		},
	}

	opts := cmp.Options{
		cmpopts.IgnoreUnexported(pgvector.Vector{}),
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
