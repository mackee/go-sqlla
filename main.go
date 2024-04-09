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

var supportedGenericsTypes = map[string]struct{}{
	"sql.Null": {},
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
		tagText := field.Tag.Value[1 : len(field.Tag.Value)-1]
		tag := reflect.StructTag(tagText)
		columnInfo := tag.Get("db")
		columnMaps := strings.Split(columnInfo, ",")
		columnName := columnMaps[0]
		isPk := false
		for _, cm := range columnMaps {
			if cm == "primarykey" {
				isPk = true
				break
			}
		}
		t := ti.TypeOf(field.Type)
		var typeName, pkgName, typeParameter string
		baseTypeName := t.String()
		nt, ok := t.(*types.Named)
		if ok {
			pkgName = nt.Obj().Pkg().Path()
			if tablePkg.Path() != pkgName {
				typeName = strings.Join([]string{nt.Obj().Pkg().Name(), nt.Obj().Name()}, ".")
			} else {
				typeName = nt.Obj().Name()
			}
			baseTypeName = typeName
			if _, ok := supportedGenericsTypes[typeName]; ok {
				tps := nt.TypeParams()
				if tps == nil {
					return nil, fmt.Errorf("toTable: has not type params: table=%s, field=%s", table.TableName, columnName)
				}
				tpsStr := make([]string, tps.Len())
				for i := 0; i < tps.Len(); i++ {
					tp := tps.At(i)
					tpsStr[i] = tp.String()
				}
				typeParameter = strings.Join(tpsStr, ",")
			} else if _, ok := supportedNonPrimitiveTypes[typeName]; !ok {
				bt := nt.Underlying()
				for _, ok := bt.Underlying().(*types.Named); ok; bt = bt.Underlying() {
				}
				baseTypeName = bt.String()
			}
		} else {
			typeName = t.String()
		}
		column := Column{
			Field:         field,
			Name:          columnName,
			IsPk:          isPk,
			typeName:      typeName,
			baseTypeName:  baseTypeName,
			PkgName:       pkgName,
			typeParameter: typeParameter,
		}
		if alt, ok := altTypeNames[baseTypeName]; ok {
			column.altTypeName = alt
		}
		table.AddColumn(column)
	}

	return table, nil
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
