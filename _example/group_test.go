package example_test

import (
	"database/sql"
	"reflect"
	"strings"
	"testing"

	example "github.com/mackee/go-sqlla/_example"
	"github.com/mackee/go-sqlla/v2"
)

var groupAllColumns = strings.Join(example.GroupAllColumns, ", ")

type queryTestCase struct {
	name  string
	query interface {
		ToSql() (string, []interface{}, error)
	}
	expect string
	args   []interface{}
}

func (q queryTestCase) assert(t *testing.T) {
	t.Helper()
	t.Run(q.name, func(t *testing.T) {
		query, args, err := q.query.ToSql()
		if err != nil {
			t.Error("unexpected error:", err)
		}
		if query != q.expect {
			t.Errorf("unexpected query: got=%s, expect=%s", query, q.expect)
		}
		if !reflect.DeepEqual(args, q.args) {
			t.Errorf("unexpected args: got=%#v, expect=%#v:", args, q.args)
		}
	})
}

type queryTestCases []queryTestCase

func (q queryTestCases) assert(t *testing.T) {
	t.Helper()
	for _, tc := range q {
		tc.assert(t)
	}
}

func TestSelectNullable(t *testing.T) {
	testCases := queryTestCases{
		{
			name:   "IS NULL",
			query:  example.NewGroupSQL().Select().SubLeaderUserIDIsNull(),
			expect: "SELECT " + groupAllColumns + " FROM `group` WHERE `sub_leader_user_id` IS NULL;",
			args:   []interface{}{},
		},
		{
			name:   "IS NOT NULL",
			query:  example.NewGroupSQL().Select().ChildGroupIDIsNotNull(),
			expect: "SELECT " + groupAllColumns + " FROM `group` WHERE `child_group_id` IS NOT NULL;",
			args:   []interface{}{},
		},
		{
			name:   "query by type parameter",
			query:  example.NewGroupSQL().Select().SubLeaderUserID(42),
			expect: "SELECT " + groupAllColumns + " FROM `group` WHERE `sub_leader_user_id` = ?;",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name:   "query by type parameter with operator",
			query:  example.NewGroupSQL().Select().SubLeaderUserID(42, sqlla.OpLess),
			expect: "SELECT " + groupAllColumns + " FROM `group` WHERE `sub_leader_user_id` < ?;",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name:   "query by type parameters multiple",
			query:  example.NewGroupSQL().Select().SubLeaderUserIDIn(42, 43, 44),
			expect: "SELECT " + groupAllColumns + " FROM `group` WHERE `sub_leader_user_id` IN(?,?,?);",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 43, Valid: true},
				sql.Null[int64]{V: 44, Valid: true},
			},
		},
	}
	testCases.assert(t)
}

func TestUpdateNullable(t *testing.T) {
	testCases := queryTestCases{
		{
			name:   "SET NOT NULL WHERE IS NULL",
			query:  example.NewGroupSQL().Update().SetSubLeaderUserID(42).WhereSubLeaderUserIDIsNull(),
			expect: "UPDATE `group` SET `sub_leader_user_id` = ? WHERE `sub_leader_user_id` IS NULL;",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name:   "SET NULL WHERE IS NULL",
			query:  example.NewGroupSQL().Update().SetSubLeaderUserIDToNull().WhereSubLeaderUserIDIsNull(),
			expect: "UPDATE `group` SET `sub_leader_user_id` = ? WHERE `sub_leader_user_id` IS NULL;",
			args:   []interface{}{sql.Null[int64]{Valid: false}},
		},
		{
			name:   "SET NULL WHERE IS NOT NULL",
			query:  example.NewGroupSQL().Update().SetChildGroupIDToNull().WhereSubLeaderUserIDIsNotNull(),
			expect: "UPDATE `group` SET `child_group_id` = ? WHERE `sub_leader_user_id` IS NOT NULL;",
			args:   []interface{}{sql.Null[int64]{Valid: false}},
		},
		{
			name:   "SET NOT NULL WHERE equal type parameters",
			query:  example.NewGroupSQL().Update().SetChildGroupID(42).WhereChildGroupID(100),
			expect: "UPDATE `group` SET `child_group_id` = ? WHERE `child_group_id` = ?;",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 100, Valid: true},
			},
		},
		{
			name:   "SET NOT NULL WHERE type parameters and operator",
			query:  example.NewGroupSQL().Update().SetChildGroupID(42).WhereChildGroupID(100, sqlla.OpGreater),
			expect: "UPDATE `group` SET `child_group_id` = ? WHERE `child_group_id` > ?;",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 100, Valid: true},
			},
		},
		{
			name:   "SET NOT NULL WHERE IN type parameters",
			query:  example.NewGroupSQL().Update().SetChildGroupID(42).WhereChildGroupIDIn(100, 101, 102),
			expect: "UPDATE `group` SET `child_group_id` = ? WHERE `child_group_id` IN(?,?,?);",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 100, Valid: true},
				sql.Null[int64]{V: 101, Valid: true},
				sql.Null[int64]{V: 102, Valid: true},
			},
		},
	}
	testCases.assert(t)
}

func TestInsertNullable(t *testing.T) {
	testCases := queryTestCases{
		{
			name:   "INSERT NULL column",
			query:  example.NewGroupSQL().Insert().ValueSubLeaderUserIDIsNull(),
			expect: "INSERT INTO `group` (`sub_leader_user_id`) VALUES(?);",
			args:   []interface{}{sql.Null[int64]{Valid: false}},
		},
		{
			name:   "INSERT with type parameter",
			query:  example.NewGroupSQL().Insert().ValueSubLeaderUserID(42),
			expect: "INSERT INTO `group` (`sub_leader_user_id`) VALUES(?);",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name: "INSERT ON DUPLICATE KEY UPDATE SET with type parameter",
			query: example.NewGroupSQL().Insert().ValueSubLeaderUserID(42).
				OnDuplicateKeyUpdate().
				ValueOnUpdateSubLeaderUserID(43),
			expect: "INSERT INTO `group` (`sub_leader_user_id`) VALUES(?) ON DUPLICATE KEY UPDATE `sub_leader_user_id` = ?;",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 43, Valid: true},
			},
		},
		{
			name: "INSERT ON DUPLICATE KEY UPDATE SET TO NULL",
			query: example.NewGroupSQL().Insert().ValueSubLeaderUserID(42).
				OnDuplicateKeyUpdate().
				ValueOnUpdateSubLeaderUserIDToNull(),
			expect: "INSERT INTO `group` (`sub_leader_user_id`) VALUES(?) ON DUPLICATE KEY UPDATE `sub_leader_user_id` = ?;",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{Valid: false},
			},
		},
	}
	testCases.assert(t)
}

func TestDeleteNullable(t *testing.T) {
	testCases := queryTestCases{
		{
			name:   "IS NULL",
			query:  example.NewGroupSQL().Delete().SubLeaderUserIDIsNull(),
			expect: "DELETE FROM `group` WHERE `sub_leader_user_id` IS NULL;",
			args:   []interface{}{},
		},
		{
			name:   "IS NOT NULL",
			query:  example.NewGroupSQL().Delete().ChildGroupIDIsNotNull(),
			expect: "DELETE FROM `group` WHERE `child_group_id` IS NOT NULL;",
			args:   []interface{}{},
		},
		{
			name:   "query by type parameter",
			query:  example.NewGroupSQL().Delete().SubLeaderUserID(42),
			expect: "DELETE FROM `group` WHERE `sub_leader_user_id` = ?;",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name:   "query by type parameter with operator",
			query:  example.NewGroupSQL().Delete().SubLeaderUserID(42, sqlla.OpLess),
			expect: "DELETE FROM `group` WHERE `sub_leader_user_id` < ?;",
			args:   []interface{}{sql.Null[int64]{V: 42, Valid: true}},
		},
		{
			name:   "query by type parameters multiple",
			query:  example.NewGroupSQL().Delete().SubLeaderUserIDIn(42, 43, 44),
			expect: "DELETE FROM `group` WHERE `sub_leader_user_id` IN(?,?,?);",
			args: []interface{}{
				sql.Null[int64]{V: 42, Valid: true},
				sql.Null[int64]{V: 43, Valid: true},
				sql.Null[int64]{V: 44, Valid: true},
			},
		},
	}
	testCases.assert(t)
}
