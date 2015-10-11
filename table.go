package sqlla

import (
	"io"

	"github.com/favclip/genbase"
)

type Table struct {
	PackageName string
	Name        string
	Columns     Columns
}

func (t *Table) AddColumn(c Column) {
	c.TableName = t.Name
	t.Columns = append(t.Columns, c)
}

func (t Table) Render(w io.Writer) error {
	return nil
}

type Columns []Column

type Column struct {
	*genbase.FieldInfo
	Name      string
	TableName string
}

func (c Column) String() string {
	return c.Name
}
