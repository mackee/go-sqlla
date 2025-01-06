//go:build !tinygo.wasm

package sqlla

import (
	"bytes"
	"embed"
	"go/format"
	"io"
	"log"
	"text/template"

	"github.com/Masterminds/goutils"
	sprig "github.com/go-task/slim-sprig"
	"github.com/pkg/errors"
	"github.com/serenize/snaker"
)

//go:embed template/* template/plugins/*
var templates embed.FS

//go:embed template/table.tmpl
var tableTmpl []byte

var tmpl = template.New("table")

func init() {
	fm := sprig.FuncMap()
	fm["untitle"] = func(s string) string {
		return goutils.Uncapitalize(s)
	}
	fm["toSnake"] = snaker.CamelToSnake
	fm["toCamel"] = snaker.SnakeToCamel
	tmpl = tmpl.Funcs(fm)

	var err error
	tmpl, err = tmpl.ParseFS(templates, "template/*.tmpl", "template/plugins/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.Parse(string(tableTmpl))
	if err != nil {
		log.Fatal(err)
	}
}

func WriteCode(w io.Writer, table *Table) error {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, table); err != nil {
		return errors.Wrapf(err, "fail to render")
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		if _, err := w.Write(buf.Bytes()); err != nil {
			return errors.Wrapf(err, "fail to write: table=%s", table.Name)
		}
		return errors.Wrapf(err, "fail to format: table=%s", table.Name)
	}
	if _, err := w.Write(bs); err != nil {
		return errors.Wrapf(err, "fail to write: table=%s", table.Name)
	}
	return nil
}
