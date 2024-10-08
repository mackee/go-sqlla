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

	"github.com/Masterminds/goutils"
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
			"Title": func(s string) string {
				return goutils.Capitalize(s)
			},
			"Untitle": func(s string) string {
				return goutils.Uncapitalize(s)
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
		if _, err := w.Write(buf.Bytes()); err != nil {
			return errors.Wrapf(err, "fail to write: table=%s", table.Name)
		}
		return errors.Wrapf(err, "fail to format: table=%s", table.Name)
	}
	_, err = w.Write(bs)
	return err
}
