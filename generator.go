package sqlla

import (
	"io"
	"strings"
	"text/template"
	"unicode"

	"github.com/serenize/snaker"
)

//go:generate go-bindata -pkg sqlla template

var templates = []string{
	"template/table.tmpl",
	"template/select.tmpl",
	"template/select_column.tmpl",
	"template/update.tmpl",
	"template/update_column.tmpl",
	"template/insert.tmpl",
	"template/insert_column.tmpl",
	"template/delete.tmpl",
	"template/delete_column.tmpl",
}

var tmpl = template.New("sqlla")

func init() {
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"Title": strings.Title,
			"Exprize": func(s string) string {
				s = strings.TrimPrefix(s, "sql.")
				s = strings.TrimPrefix(s, "time.")
				s = strings.TrimPrefix(s, "mysql.")

				return s
			},
			"Untitle": func(s string) string {
				s0 := rune(s[0])
				if !unicode.IsUpper(s0) {
					return s
				}
				s0l := unicode.ToLower(rune(s[0]))
				return string(s0l) + s[1:]
			},
			"toLower": strings.ToLower,
			"toSnake": snaker.CamelToSnake,
			"toCamel": snaker.SnakeToCamel,
		},
	)
	for _, filename := range templates {
		data, err := Asset(filename)
		tmpl, err = tmpl.Parse(string(data))
		if err != nil {
			panic(err)
		}
	}
}

func WriteCode(w io.Writer, table *Table) error {
	return tmpl.Execute(w, table)
}
