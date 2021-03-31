package sqlla

import (
	"bytes"
	"go/format"
	"io"
	"text/template"

	"github.com/pkg/errors"
)

var tmpl = template.New("table")

func WriteCode(w io.Writer, table *Table) error {
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, table)
	if err != nil {
		return errors.Wrapf(err, "fail to render")
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrapf(err, "fail to format")
	}
	_, err = w.Write(bs)
	return err
}
