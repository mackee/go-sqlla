package sqlla

import (
	"io"
	"strings"

	"github.com/favclip/genbase"
)

type Table struct {
	PackageName           string
	StructName            string
	Name                  string
	Columns               Columns
	PkColumn              *Column
	additionalPackagesMap map[string]struct{}
}

func (t *Table) AddColumn(c Column) {
	if t.additionalPackagesMap == nil {
		t.additionalPackagesMap = make(map[string]struct{})
	}
	c.TableName = t.Name
	if c.IsPk {
		t.PkColumn = &c
		t.additionalPackagesMap["database/sql"] = struct{}{}
	}
	switch {
	case strings.HasPrefix(c.TypeName(), "sql.Null"):
		t.additionalPackagesMap["database/sql"] = struct{}{}
	case c.TypeName() == "time.Time":
		t.additionalPackagesMap["time"] = struct{}{}
	case c.TypeName() == "mysql.NullTime":
		t.additionalPackagesMap["github.com/go-sql-driver/mysql"] = struct{}{}
	}
	t.Columns = append(t.Columns, c)
}

func (t *Table) AdditionalPackages() []string {
	packages := make([]string, 0, len(t.additionalPackagesMap))
	for pkg, _ := range t.additionalPackagesMap {
		packages = append(packages, pkg)
	}
	return packages
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
