package genbase

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/types"
	"strings"

	_ "golang.org/x/tools/go/gcimporter"
)

var (
	ErrNotStructType = errors.New("type is not ast.StructType")
)

type Parser struct {
	SkipSemanticsCheck bool
}

type PackageInfo struct {
	Dir   string
	Files FileInfos
	Types *types.Package
}

type FileInfo ast.File
type FileInfos []*FileInfo

// try http://goast.yuroyoro.net/ with http://play.golang.org/p/ruqMMsbDaw
type TypeInfo struct {
	FileInfo *FileInfo
	GenDecl  *ast.GenDecl
	TypeSpec *ast.TypeSpec
	Comment  *ast.Comment
}
type TypeInfos []*TypeInfo

type StructTypeInfo ast.StructType

type FieldInfo ast.Field
type FieldInfos []*FieldInfo

func (p *Parser) ParsePackageDir(directory string) (*PackageInfo, error) {
	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot process directory %s: %s", directory, err)
	}
	names := make([]string, 0)
	names = append(names, pkg.GoFiles...)
	names = append(names, pkg.CgoFiles...)
	names = append(names, pkg.SFiles...)
	names = pathJoinAll(directory, names...)
	return p.parsePackage(directory, names)
}

func (p *Parser) ParsePackageFiles(fileNames []string) (*PackageInfo, error) {
	return p.parsePackage(".", fileNames)
}

func (p *Parser) parsePackage(directory string, fileNames []string) (*PackageInfo, error) {
	var files FileInfos
	pkg := &PackageInfo{}
	fs := token.NewFileSet()
	for _, fileName := range fileNames {
		if !strings.HasSuffix(fileName, ".go") {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, fileName, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parsing package: %s: %s", fileName, err)
		}
		files = append(files, (*FileInfo)(parsedFile))
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("%s: no buildable Go files", directory)
	}
	pkg.Files = files
	pkg.Dir = directory

	// resolve types
	config := types.Config{
		FakeImportC:              true,
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
	}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	typesPkg, err := config.Check(pkg.Dir, fs, files.AstFiles(), info)
	if p.SkipSemanticsCheck && err != nil {
		return pkg, nil
	} else if err != nil {
		return nil, err
	}
	pkg.Types = typesPkg

	return pkg, nil
}

func (pkg *PackageInfo) TypeInfos() TypeInfos {
	var types TypeInfos
	for _, file := range pkg.Files {
		if file == nil {
			continue
		}
		ast.Inspect(file.AstFile(), func(node ast.Node) bool {
			decl, ok := node.(*ast.GenDecl)
			if !ok || decl.Tok != token.TYPE {
				return true
			}
			found := false
			for _, spec := range decl.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				types = append(types, &TypeInfo{
					FileInfo: file,
					GenDecl:  decl,
					TypeSpec: ts,
				})
				found = true
			}
			return !found
		})
	}
	return types
}

func (pkg *PackageInfo) CollectTaggedTypeInfos(tag string) TypeInfos {
	ret := TypeInfos{}

	types := pkg.TypeInfos()

	for _, t := range types {
		if c := findAnnotation(t.Doc(), tag); c != nil {
			t.Comment = c
			ret = append(ret, t)
		}
	}

	return ret
}

func (pkg *PackageInfo) CollectTypeInfos(typeNames []string) TypeInfos {
	ret := TypeInfos{}

	types := pkg.TypeInfos()

outer:
	for _, t := range types {
		for _, name := range typeNames {
			if t.Name() == name {
				ret = append(ret, t)
				continue outer
			}
		}
	}

	return ret
}

func (pkg *PackageInfo) Name() string {
	return pkg.Files[0].Name.Name
}

func (file *FileInfo) AstFile() *ast.File {
	return (*ast.File)(file)
}

func (files FileInfos) AstFiles() []*ast.File {
	astFiles := make([]*ast.File, len(files))
	for i, file := range files {
		astFiles[i] = file.AstFile()
	}
	return astFiles
}

func (file *FileInfo) FindImportSpecByIdent(packageIdent string) *ast.ImportSpec {
	for _, imp := range file.Imports {
		if imp.Name != nil && imp.Name.Name == packageIdent {
			// import foo "foobar"
			return imp
		} else if strings.HasSuffix(imp.Path.Value, fmt.Sprintf(`/%s"`, packageIdent)) {
			// import "favclip/foo"
			return imp
		} else if imp.Path.Value == fmt.Sprintf(`"%s"`, packageIdent) {
			// import "foo"
			return imp
		}
	}
	return nil
}

func (t *TypeInfo) StructType() (*StructTypeInfo, error) {
	structType, ok := interface{}(t.TypeSpec.Type).(*ast.StructType)
	if !ok {
		return nil, ErrNotStructType
	}

	return (*StructTypeInfo)(structType), nil
}

func (t *TypeInfo) Name() string {
	return t.TypeSpec.Name.Name
}

func (t *TypeInfo) Doc() *ast.CommentGroup {
	if t.TypeSpec.Doc != nil {
		return t.TypeSpec.Doc
	}
	if t.GenDecl.Doc != nil {
		return t.GenDecl.Doc
	}
	return nil
}

func (st *StructTypeInfo) AstStructType() *ast.StructType {
	return (*ast.StructType)(st)
}

func (st *StructTypeInfo) FieldInfos() FieldInfos {
	var fields FieldInfos
	for _, field := range st.AstStructType().Fields.List {
		fields = append(fields, (*FieldInfo)(field))
	}

	return fields
}

func (f *FieldInfo) TypeName() string {
	typeName, err := ExprToTypeName(f.Type)
	if err != nil {
		return fmt.Sprintf("!!%s!!", err.Error())
	}
	return typeName
}

func (f *FieldInfo) IsPtr() bool {
	_, ok := f.Type.(*ast.StarExpr)
	return ok
}

func (f *FieldInfo) IsArray() bool {
	_, ok := f.Type.(*ast.ArrayType)
	return ok
}

func (f *FieldInfo) IsPtrArray() bool {
	star, ok := f.Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	_, ok = star.X.(*ast.ArrayType)
	return ok
}

func (f *FieldInfo) IsArrayPtr() bool {
	array, ok := f.Type.(*ast.ArrayType)
	if !ok {
		return false
	}
	_, ok = array.Elt.(*ast.StarExpr)
	return ok
}

func (f *FieldInfo) IsPtrArrayPtr() bool {
	star, ok := f.Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	array, ok := star.X.(*ast.ArrayType)
	if !ok {
		return false
	}
	_, ok = array.Elt.(*ast.StarExpr)
	return ok
}

func (field *FieldInfo) IsInt64() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "int64"
}

func (field *FieldInfo) IsInt() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "int"
}

func (field *FieldInfo) IsString() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "string"
}

func (field *FieldInfo) IsFloat32() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "float32"
}

func (field *FieldInfo) IsFloat64() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "float64"
}

func (field *FieldInfo) IsNumber() bool {
	return field.IsInt() || field.IsInt64() || field.IsFloat32() || field.IsFloat64()
}

func (field *FieldInfo) IsBool() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "bool"
}

func (field *FieldInfo) IsTime() bool {
	typeName, err := ExprToBaseTypeName(field.Type)
	if err != nil {
		return false
	}
	return typeName == "time.Time"
}
