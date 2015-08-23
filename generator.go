package sqlla

import (
	"io"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

//go:generate go-bindata -pkg sqlla template

var tmpl = template.New("sqlla")

func init() {
	tableTmpl, _ := Asset("template/table.tmpl")
	columnTmpl, _ := Asset("template/column.tmpl")

	tmpl = tmpl.Funcs(
		template.FuncMap{
			"Title":   strings.Title,
			"toLower": strings.ToLower,
			"toSnake": snaker.CamelToSnake,
			"toCamel": snaker.SnakeToCamel,
		},
	)
	var err error
	tmpl, err = tmpl.Parse(string(tableTmpl))
	if err != nil {
		panic(err)
	}
	tmpl, err = tmpl.Parse(string(columnTmpl))
	if err != nil {
		panic(err)
	}
}

func WriteCode(w io.Writer, table *Table) error {
	return tmpl.Execute(w, table)
}
