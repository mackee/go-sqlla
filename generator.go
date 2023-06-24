//go:build !tinygo.wasm

package sqlla

import (
	"bytes"
	"embed"
	"go/format"
	"io"
	"log"
	"strings"
	"text/template"
	"unicode"

	"github.com/pkg/errors"
	"github.com/serenize/snaker"
)

//go:embed template/*
var templates embed.FS

//go:embed template/table.tmpl
var tableTmpl []byte

var tmpl = template.New("table")

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
	var err error
	tmpl, err = tmpl.ParseFS(templates, "template/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(string(tableTmpl))
	if err != nil {
		log.Fatal(err)
	}
}

func WriteCode(w io.Writer, table *Table) error {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, table)
	if err != nil {
		return errors.Wrapf(err, "fail to render")
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrapf(err, "fail to format: table=%s", table.Name)
	}
	_, err = w.Write(bs)
	return err
}
