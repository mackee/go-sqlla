package sqlla

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/favclip/genbase"
)

func Run(from string) {
	p := &genbase.Parser{}
	dir := filepath.Dir(from)
	pinfo, err := p.ParsePackageFiles([]string{from})
	if err != nil {
		panic(err)
	}

	typeInfos := pinfo.CollectTaggedTypeInfos("+table:")
	for _, typeInfo := range typeInfos {
		table, err := toTable(typeInfo)
		if err != nil {
			panic(err)
		}
		filename := filepath.Join(dir, table.Name+"_auto.go")
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		WriteCode(f, table)
	}
}

func toTable(typeInfo *genbase.TypeInfo) (*Table, error) {
	table := new(Table)
	table.PackageName = typeInfo.FileInfo.Name.Name

	comment := typeInfo.AnnotatedComment.Text
	tableName := strings.TrimPrefix(comment, "//+table: ")
	table.Name = tableName

	table.StructName = typeInfo.Name()

	structType, err := typeInfo.StructType()
	if err != nil {
		return nil, err
	}

	fieldInfos := structType.FieldInfos()
	for _, fieldInfo := range fieldInfos {
		tagText := fieldInfo.Tag.Value[1 : len(fieldInfo.Tag.Value)-1]
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
		column := Column{FieldInfo: fieldInfo, Name: columnName, IsPk: isPk}
		table.AddColumn(column)
	}

	return table, nil
}
