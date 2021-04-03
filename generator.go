package sqlla

import (
	"fmt"
	"io"
	"runtime"
	"text/template"
)

var tmpl = template.New("table")

func WriteCode(w io.Writer, table *Table) error {
	return fmt.Errorf("must build by go version of 1.16 or higher. this built by %s", runtime.Version())
}
