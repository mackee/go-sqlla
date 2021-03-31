package sqlla

import (
	"go/ast"
	"io"
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
	}
	if c.PkgName != "" {
		t.additionalPackagesMap[c.PkgName] = struct{}{}
	}
	t.Columns = append(t.Columns, c)
}

func (t *Table) AdditionalPackages() []string {
	packages := make([]string, 0, len(t.additionalPackagesMap))
	for pkg := range t.additionalPackagesMap {
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
	Field        *ast.Field
	Name         string
	TypeName     string
	PkgName      string
	BaseTypeName string
	TableName    string
	IsPk         bool
}

func (c Column) String() string {
	return c.Name
}

func (c Column) FieldName() string {
	if len(c.Field.Names) > 0 {
		return c.Field.Names[0].Name
	}
	return ""
}
