//go:build !tinygo.wasm

package sqlla

import (
	"go/ast"
	"go/types"
	"io"
	"strings"

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

type Columns []Column

type Column struct {
	Field         *ast.Field
	Name          string
	MethodName    string
	typeName      string
	PkgName       string
	baseTypeName  string
	altTypeName   string
	typeParameter string
	TableName     string
	IsPk          bool
}

func (c Column) HasUnderlyingType() bool {
	return c.baseTypeName != c.typeName
}

func (c Column) TypeName() string {
	tn := c.typeName
	if c.typeParameter != "" {
		return tn + "[" + c.typeParameter + "]"
	}
	return tn
}

func (c Column) BaseTypeName() string {
	if c.typeParameter != "" {
		return c.baseTypeName + "[" + c.typeParameter + "]"
	}
	return c.baseTypeName
}

func (c Column) AltTypeName() string {
	if c.altTypeName == "" {
		return ""
	}
	if c.typeParameter != "" {
		return c.altTypeName + "[" + c.typeParameter + "]"
	}
	return c.altTypeName
}

func (c Column) typeNameForExpr() string {
	tname := c.BaseTypeName()
	if atn := c.AltTypeName(); atn != "" {
		tname = atn
	}
	tname = c.exprize(tname)
	tname = goutils.Capitalize(tname)
	return tname
}

func (c Column) ExprTypeName() string {
	return "Expr" + c.typeNameForExpr()
}

func (c Column) exprize(s string) string {
	s = strings.TrimPrefix(s, "sql.")
	s = strings.TrimPrefix(s, "time.")
	s = strings.TrimPrefix(s, "mysql.")
	return s
}

func (c Column) ExprMultiTypeName() string {
	return "ExprMulti" + c.typeNameForExpr()
}

func (c Column) ExprValueIdentifier() string {
	if c.typeName != c.baseTypeName {
		return c.baseTypeName + "(v)"
	}
	return "v"
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
