package sqlla_test

import (
	"reflect"
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
					Name:    "count",
					Outpath: "user_count.gen.go",
					Args:    map[string]string{},
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
					Name:    "table",
					Outpath: "user_table.gen.go",
					Args: map[string]string{
						"getcolumns":    "id,name",
						"listincolumns": "id,name",
						"listcolumns":   "age",
					},
				},
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sqlla.ParsePluginsByComments(tt.comments)
			if err != nil {
				t.Errorf("parsePlguinByComments() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePlguinByComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
