// +build go1.16

package sqlla

import (
	"embed"
	"log"
	"strings"
	"text/template"
	"unicode"

	"github.com/serenize/snaker"
)

//go:embed template/*
var templates embed.FS

//go:embed template/table.tmpl
var tableTmpl []byte

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
