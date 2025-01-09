//go:build !tinygo.wasm

package sqlla

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Options struct {
	Version bool     `help:"show this version"`
	From    []string `help:"source file path" env:"GOFILE" arg:"" required:""`
	Ext     string   `help:"file extension" env:"SQLLA_GENERATE_FILE_EXT" default:".gen.go"`
	Plugins []string `help:"additional plugin files. allow asterisk(*) and multiple values" name:"plugins" env:"SQLLA_TEMPLATE_FILES"`
	Dialect string   `help:"SQL Dialect" env:"SQLLA_DIALECT" enum:"mysql,sqlite,postgresql" default:"mysql"`
}

func Run(opts Options) error {
	dialect, err := NewDialect(opts.Dialect)
	if err != nil {
		return fmt.Errorf("failed to NewDialect: %w", err)
	}
	g, err := NewGenerator(dialect, opts.Plugins...)
	if err != nil {
		return fmt.Errorf("failed to NewGenerator: %w", err)
	}

	for _, from := range opts.From {
		if err := run(from, opts.Ext, g); err != nil {
			return fmt.Errorf("failed to run: %w", err)
		}
	}
	return nil
}

func run(from string, ext string, g *Generator) error {
	fullpath, err := filepath.Abs(from)
	if err != nil {
		return fmt.Errorf("failed to filepath.Abs: %w", err)
	}
	dir := filepath.Dir(fullpath)

	pkgs, err := toPackages(dir)
	if err != nil {
		return fmt.Errorf("failed to toPackages: %w", err)
	}
	for _, pkg := range pkgs {
		files := pkg.Syntax
		for _, f := range files {
			for _, decl := range f.Decls {
				table, err := declToTable(pkg, decl, fullpath)
				if err != nil {
					if errors.Is(err, errNotTargetDecl) {
						continue
					}
					return fmt.Errorf("error declToTable: %w", err)
				}
				filename := filepath.Join(dir, table.TableName+ext)
				f, err := os.Create(filename)
				if err != nil {
					return fmt.Errorf("error create: filename=%s: %w", filename, err)
				}
				if err := g.WriteCode(f, table); err != nil {
					return fmt.Errorf("error WriteCode: filename=%s: %w", filename, err)
				}
				if err := f.Close(); err != nil {
					return fmt.Errorf("error close: filename=%s: %w", filename, err)
				}
				if err := table.Plugins.WriteCode(g, table.PackageName); err != nil {
					return fmt.Errorf("error Plugins.WriteCode: %w", err)
				}
			}
		}
	}
	return nil
}

func toPackages(dir string) ([]*packages.Package, error) {
	conf := &packages.Config{
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(conf, dir)
	if err != nil {
		return nil, fmt.Errorf("error toPackages: %w", err)
	}
	return pkgs, nil
}

var errNotTargetDecl = fmt.Errorf("not target decl")

func declToTable(pkg *packages.Package, decl ast.Decl, fullpath string) (*Table, error) {
	pos := pkg.Fset.Position(decl.Pos())
	if pos.Filename != fullpath {
		return nil, errNotTargetDecl
	}
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil, errNotTargetDecl
	}
	if genDecl.Doc == nil {
		return nil, errNotTargetDecl
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
		return nil, errNotTargetDecl
	}
	table, err := toTable(pkg.Types, annotationComment, genDecl, pkg.TypesInfo)
	if err != nil {
		return nil, fmt.Errorf("error toTable: %w", err)
	}
	return table, nil
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
	table.additionalPackagesMap = make(map[string]struct{})

	table.TableName = trimAnnotation(annotationComment)
	qualifier := types.RelativeTo(tablePkg)

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
	comments := make([]string, 0, len(gd.Doc.List))
	for _, comment := range gd.Doc.List {
		comments = append(comments, comment.Text)
	}
	plugins, err := parsePluginsByComments(comments)
	if err != nil {
		return nil, fmt.Errorf("toTable: error parsePluginsByComments: table=%s, err=%w", table.TableName, err)
	}
	table.SetPlugins(plugins)

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
				tas := nt.TypeArgs()
				if tas == nil {
					return nil, fmt.Errorf("toTable: has not type params: table=%s, field=%s", table.TableName, columnName)
				}
				tpsStr := make([]string, tas.Len())
				for i := 0; i < tas.Len(); i++ {
					ta := tas.At(i)
					switch ata := ta.(type) {
					case *types.Named:
						tn := ata.Obj()
						tpsStr[i] = tn.Id()
						if qualifier(tn.Pkg()) != "" {
							tpsStr[i] = tn.Pkg().Name() + "." + tn.Id()
							table.additionalPackagesMap[tn.Pkg().Path()] = struct{}{}
						}
					case *types.Basic:
						tpsStr[i] = ata.Name()
					default:
						return nil, fmt.Errorf("toTable: unsupported type param: table=%s, field=%s, type=%s", table.TableName, columnName, ta.String())
					}
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
