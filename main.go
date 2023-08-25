//go:build !tinygo.wasm

package sqlla

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

func Run(from, ext string) {
	fullpath, err := filepath.Abs(from)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(fullpath)

	conf := &packages.Config{
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(conf, dir)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		files := pkg.Syntax
		for _, f := range files {
			for _, decl := range f.Decls {
				pos := pkg.Fset.Position(decl.Pos())
				if pos.Filename != fullpath {
					continue
				}
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				if genDecl.Doc == nil {
					continue
				}
				var hasAnnotation bool
				var annotationComment string
				for _, comment := range genDecl.Doc.List {
					if trimmed := trimAnnotation(comment.Text); trimmed != comment.Text {
						hasAnnotation = true
						annotationComment = comment.Text
					}
				}
				if !hasAnnotation {
					continue
				}
				table, err := toTable(pkg.Types, annotationComment, genDecl, pkg.TypesInfo)
				if err != nil {
					panic(err)
				}
				filename := filepath.Join(dir, table.TableName+ext)
				f, err := os.Create(filename)
				if err != nil {
					panic(err)
				}
				if err := WriteCode(f, table); err != nil {
					panic(err)
				}
			}
		}
	}
}

var supportedNonPrimitiveTypes = map[string]struct{}{
	"time.Time":       {},
	"mysql.NullTime":  {},
	"sql.NullTime":    {},
	"sql.NullInt64":   {},
	"sql.NullString":  {},
	"sql.NullFloat64": {},
	"sql.NullBool":    {},
}

var altTypeNames = map[string]string{
	"[]byte":         "Bytes",
	"mysql.NullTime": "MysqlNullTime",
}

func toTable(tablePkg *types.Package, annotationComment string, gd *ast.GenDecl, ti *types.Info) (*Table, error) {
	table := new(Table)
	table.Package = tablePkg
	table.PackageName = tablePkg.Name()

	table.TableName = trimAnnotation(annotationComment)

	spec := gd.Specs[0]
	ts, ok := spec.(*ast.TypeSpec)
	if !ok {
		return nil, fmt.Errorf("toTable: not type spec: table=%s", table.TableName)
	}
	structType, ok := ts.Type.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("toTable: not struct type: table=%s", table.TableName)
	}
	table.StructName = ts.Name.Name

	if isV2Annotation(annotationComment) {
		table.Name = table.StructName
	} else {
		table.Name = table.TableName
	}

	for _, field := range structType.Fields.List {
		columns := toColumns(toColumnsInput{
			field: field,
			ti:    ti,
			pkg:   tablePkg,
		})
		table.AddColumns(columns)
	}

	return table, nil
}

type dbTag struct {
	columnName string
	attrs      []string
}

func parseDBTag(tagText string) *dbTag {
	tag := reflect.StructTag(tagText)
	columnInfo := tag.Get("db")
	if columnInfo == "" {
		return nil
	}

	attrs := strings.Split(columnInfo, ",")
	columnName := attrs[0]

	return &dbTag{
		columnName: columnName,
		attrs:      attrs,
	}
}

func (d *dbTag) nested() bool {
	for _, attr := range d.attrs[1:] {
		if attr == "nested" {
			return true
		}
	}
	return false
}

func (d *dbTag) isPrimaryKey() bool {
	for _, attr := range d.attrs[1:] {
		if attr == "primarykey" {
			return true
		}
	}
	return false
}

type toColumnsInput struct {
	field *ast.Field
	ti    *types.Info
	pkg   *types.Package
}

func toColumns(input toColumnsInput) []Column {
	tagText := input.field.Tag.Value[1 : len(input.field.Tag.Value)-1]
	dt := parseDBTag(tagText)
	if dt == nil {
		return nil
	}
	if dt.nested() {
		return toNestedColumns(toNestedColumnsInput{
			field:  input.field,
			ti:     input.ti,
			pkg:    input.pkg,
			prefix: dt.columnName,
		})
	}

	fieldName := input.field.Names[0].Name
	t := input.ti.TypeOf(input.field.Type)
	return []Column{toColumn(toColumnInput{
		fieldName: fieldName,
		t:         t,
		pkg:       input.pkg,
		dbTag:     dt,
	})}
}

type toColumnInput struct {
	fieldName string
	t         types.Type
	pkg       *types.Package
	dbTag     *dbTag
}

func toColumn(input toColumnInput) Column {
	isPk := input.dbTag.isPrimaryKey()
	t := input.t
	var typeName, pkgName string
	baseTypeName := t.String()
	nt, ok := t.(*types.Named)
	if ok {
		pkgName = nt.Obj().Pkg().Path()
		if input.pkg.Path() != pkgName {
			typeName = strings.Join([]string{nt.Obj().Pkg().Name(), nt.Obj().Name()}, ".")
		} else {
			typeName = nt.Obj().Name()
		}
		baseTypeName = typeName
		if _, ok := supportedNonPrimitiveTypes[typeName]; !ok {
			bt := nt.Underlying()
			for _, ok := bt.Underlying().(*types.Named); ok; bt = bt.Underlying() {
			}
			baseTypeName = bt.String()
		}
	} else {
		typeName = t.String()
	}
	column := Column{
		FieldName:    input.fieldName,
		Name:         input.dbTag.columnName,
		IsPk:         isPk,
		TypeName:     typeName,
		BaseTypeName: baseTypeName,
		PkgName:      pkgName,
	}
	if alt, ok := altTypeNames[baseTypeName]; ok {
		column.AltTypeName = alt
	}
	return column
}

type toNestedColumnsInput struct {
	field  *ast.Field
	ti     *types.Info
	pkg    *types.Package
	prefix string
}

func toNestedColumns(input toNestedColumnsInput) []Column {
	t := input.ti.TypeOf(input.field.Type)
	tn, ok := t.(*types.Named)
	if !ok {
		return nil
	}
	ut := tn.Underlying()
	st, ok := ut.(*types.Struct)
	if !ok {
		return nil
	}
	columns := make([]Column, 0, st.NumFields())
	fieldName := input.field.Names[0].Name
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		tagText := st.Tag(i)
		dt := parseDBTag(tagText)
		if dt == nil {
			continue
		}

		column := toColumn(toColumnInput{
			fieldName: field.Name(),
			t:         field.Type(),
			pkg:       input.pkg,
			dbTag:     dt,
		})
		column.FieldName = strings.Join([]string{fieldName, column.FieldName}, ".")
		column.Name = input.prefix + column.Name
		column.IsPk = false
		columns = append(columns, column)
	}

	return columns
}

func trimAnnotation(comment string) string {
	prefixes := []string{"//+table: ", "// +table: ", "//sqlla:table "}
	for _, prefix := range prefixes {
		if trimmed := strings.TrimPrefix(comment, prefix); trimmed != comment {
			return trimmed
		}
	}
	return comment
}

func isV2Annotation(comment string) bool {
	return strings.HasPrefix(comment, "//sqlla:table ")
}
