//go:build !tinygo.wasm

package sqlla

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io"
	"text/template"

	"github.com/Masterminds/goutils"
	"github.com/gertd/go-pluralize"
	sprig "github.com/go-task/slim-sprig"
	"github.com/pkg/errors"
	"github.com/serenize/snaker"
	"golang.org/x/tools/imports"
)

//go:embed template/* template/plugins/*
var templates embed.FS

//go:embed template/table.tmpl
var tableTmpl []byte

type Generator struct {
	tmpl *template.Template
}

func NewGenerator(dialect Dialect, additionals ...string) (*Generator, error) {
	tmpl := template.New("table")

	fm := sprig.FuncMap()
	fm["untitle"] = func(s string) string {
		return goutils.Uncapitalize(s)
	}
	fm["toSnake"] = snaker.CamelToSnake
	fm["toCamel"] = snaker.SnakeToCamel
	pc := pluralize.NewClient()
	fm["pluralize"] = pc.Plural
	fm["singular"] = pc.Singular
	fm["cquote"] = dialect.CQuote
	fm["cquoteby"] = dialect.CQuoteBy
	tmpl = tmpl.Funcs(fm)

	tmpl, err := tmpl.ParseFS(templates, "template/*.tmpl", "template/plugins/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	if len(additionals) > 0 {
		for _, add := range additionals {
			tmpl, err = tmpl.ParseGlob(add)
			if err != nil {
				return nil, fmt.Errorf("failed to parse additional plugins: path=%s: %w", add, err)
			}
		}
	}
	tmpl, err = tmpl.Parse(string(tableTmpl))
	if err != nil {
		return nil, fmt.Errorf("failed to parse table template: %w", err)
	}
	return &Generator{tmpl: tmpl}, nil
}

func (g *Generator) WriteCode(w io.Writer, table *Table) error {
	buf := &bytes.Buffer{}
	if err := g.tmpl.Execute(buf, table); err != nil {
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

func (g *Generator) WriteCodeByPlugin(w io.Writer, tmplName string, p *Plugin) error {
	ptmpl := g.tmpl.Lookup(fmt.Sprintf("plugin.%s", tmplName))
	if ptmpl == nil {
		return fmt.Errorf("template not found: template=%s", tmplName)
	}
	if err := ptmpl.Execute(w, p); err != nil {
		return fmt.Errorf("fail to render: %w", err)
	}
	return nil
}

func (g *Generator) Format(input []byte, filename string) ([]byte, error) {
	out, err := imports.Process(filename, input, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to format: %w", err)
	}

	return out, nil
}
