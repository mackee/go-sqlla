package sqlla

import (
	"io"

	"github.com/favclip/genbase"
)

type Table struct {
	PackageName string
	StructName  string
	Name        string
	Columns     Columns
	PkColumn    *Column
}

func (t *Table) AddColumn(c Column) {
	c.TableName = t.Name
	if c.IsPk {
		t.PkColumn = &c
	}
	t.Columns = append(t.Columns, c)
}

func (t *Table) HasPk() bool {
	return t.PkColumn != nil
}

func (t Table) Render(w io.Writer) error {
	return nil
}

type Columns []Column

type Column struct {
	*genbase.FieldInfo
	Name      string
	TableName string
	IsPk      bool
}

func (c Column) String() string {
	return c.Name
}

func (c Column) FieldName() string {
	if len(c.Names) > 0 {
		return c.Names[0].Name
	}
	return ""
}
