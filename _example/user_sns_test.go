package example

import (
	"reflect"
	"testing"
)

func TestSelectUserSNS(t *testing.T) {
	query, args, err := NewUserSNSSQL().Select().SNSType("GITHUB").ToSql()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if query != "SELECT `id`, `sns_type`, `created_at`, `updated_at` FROM user_sns WHERE `sns_type` = ?;" {
		t.Error("unexpected query:", query)
	}
	if !reflect.DeepEqual(args, []interface{}{"GITHUB"}) {
		t.Error("unexpected args:", args)
	}
}
