//go:build !tinygo.wasm

package sqlla

import (
	"go/types"
	"io"

	"github.com/Masterminds/goutils"
	"github.com/serenize/snaker"
)

type Table struct {
	Package               *types.Package
	PackageName           string
	Name                  string
	StructName            string
	TableName             string
	Columns               Columns
	PkColumn              *Column
	additionalPackagesMap map[string]struct{}
}

func (t *Table) NamingIsStructName() bool {
	return t.Name == t.StructName
}

func (t *Table) AddColumn(c Column) {
	if t.additionalPackagesMap == nil {
		t.additionalPackagesMap = make(map[string]struct{})
	}
	c.TableName = t.Name
	if t.NamingIsStructName() {
		c.MethodName = c.FieldName()
	} else {
		c.MethodName = goutils.Capitalize(snaker.SnakeToCamel(c.Name))
	}
	if c.IsPk {
		t.PkColumn = &c
	}
	if c.PkgName != "" {
		if t.Package.Path() != c.PkgName {
			t.additionalPackagesMap[c.PkgName] = struct{}{}
		}
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
