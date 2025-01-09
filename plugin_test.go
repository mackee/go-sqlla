package sqlla_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/mackee/go-sqlla/v2"
)

func Test_parsePluginsByComments(t *testing.T) {
	tc := []struct {
		name     string
		comments []string
		want     sqlla.Plugins
	}{
		{
			name: "basic",
			comments: []string{
				"//sqlla:plugin count outpath=user_count.gen.go",
			},
			want: sqlla.Plugins{
				{
					Name: "count",
					Args: map[string]string{
						"outpath": "user_count.gen.go",
					},
				},
			},
		},
		{
			name: "multiargs",
			comments: []string{
				"//sqlla:plugin table outpath=user_table.gen.go getcolumns=id,name listincolumns=id,name listcolumns=age",
			},
			want: sqlla.Plugins{
				{
					Name: "table",
					Args: map[string]string{
						"outpath":       "user_table.gen.go",
						"getcolumns":    "id,name",
						"listincolumns": "id,name",
						"listcolumns":   "age",
					},
				},
			},
		},
	}
	prettyPrint := func(v any) string {
		bs := &strings.Builder{}
		enc := json.NewEncoder(bs)
		enc.SetIndent("", "  ")
		enc.Encode(v)
		return bs.String()
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sqlla.ParsePluginsByComments(tt.comments)
			if err != nil {
				t.Errorf("parsePlguinByComments() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePlguinByComments() = %s, want %s", prettyPrint(got), prettyPrint(tt.want))
			}
		})
	}
}
