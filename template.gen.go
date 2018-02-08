package sqlla

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets88a39ad92bb2c8db797af1c5e70ac4d41dd78298 = "{{ define \"SelectColumn\" }}{{ $smallTableName := .TableName | toCamel | Untitle }}\nfunc (q {{ $smallTableName }}SelectSQL) {{ .Name | toCamel | Title }}(v {{ .TypeName }}, exprs ...sqlla.Operator) {{ $smallTableName }}SelectSQL {\n\tvar op sqlla.Operator\n\tif len(exprs) == 0 {\n\t\top = sqlla.OpEqual\n\t} else {\n\t\top = exprs[0]\n\t}\n\n\twhere := sqlla.Expr{{ .TypeName | Exprize | Title }}{Value: v, Op: op, Column: \"{{ .Name }}\"}\n\tq.where = append(q.where, where)\n\treturn q\n}\n\nfunc (q {{ $smallTableName }}SelectSQL) {{ .Name | toCamel | Title }}In(vs ...{{ .TypeName }}) {{ $smallTableName }}SelectSQL {\n\twhere := sqlla.ExprMulti{{ .TypeName | Exprize | Title }}{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: \"{{ .Name }}\"}\n\tq.where = append(q.where, where)\n\treturn q\n}\n\n{{ if .IsPk -}}\nfunc (q {{ $smallTableName }}SelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) {{ $smallTableName }}SelectSQL {\n\tv := {{ .TypeName }}(pk)\n\treturn q.{{ .Name | toCamel | Title }}(v, exprs...)\n}\n{{- end }}\n\nfunc (q {{ $smallTableName }}SelectSQL) OrderBy{{ .Name | toCamel }}(order sqlla.Order) {{ $smallTableName }}SelectSQL {\n\tq.order = \" ORDER BY {{ .Name }}\"\n\tif order == sqlla.Asc {\n\t\tq.order += \" ASC\"\n\t} else {\n\t\tq.order += \" DESC\"\n\t}\n\n\treturn q\n}\n{{ end }}\n"
var _Assets6b103fa3a9e63224b4df28a8ed3efae32053f5bc = "package {{ .PackageName }}\n\nimport (\n\t\"strings\"\n\t\"strconv\"\n\n\t\"database/sql\"\n\t{{ range .AdditionalPackages -}}\n\t\"{{ . }}\"\n\t{{ end }}\n\t\"github.com/mackee/go-sqlla\"\n)\n{{ $camelName := .Name | toCamel | Untitle }}\ntype {{ $camelName }}SQL struct {\n\twhere sqlla.Where\n}\n\nfunc New{{ .Name | toCamel | Title }}SQL() {{ $camelName }}SQL {\n\tq := {{ $camelName }}SQL{}\n\treturn q\n}\n\n{{ template \"Select\" . }}\n{{ template \"Update\" . }}\n{{ template \"Insert\" . }}\n{{ template \"Delete\" . }}\n"
var _Assets83c5f0ba25bc78b44a12ea398292f04e3a78d0a0 = "{{ define \"UpdateColumn\" }}{{ $smallTableName := .TableName | toCamel | Untitle }}\nfunc (q {{ $smallTableName }}UpdateSQL) Set{{ .Name | toCamel | Title }}(v {{ .TypeName }}) {{ $smallTableName }}UpdateSQL {\n\tq.setMap[\"{{ .Name }}\"] = v\n\treturn q\n}\n\nfunc (q {{ $smallTableName }}UpdateSQL) Where{{ .Name | toCamel | Title }}(v {{ .TypeName }}, exprs ...sqlla.Operator) {{ $smallTableName }}UpdateSQL {\n\tvar op sqlla.Operator\n\tif len(exprs) == 0 {\n\t\top = sqlla.OpEqual\n\t} else {\n\t\top = exprs[0]\n\t}\n\n\twhere := sqlla.Expr{{ .TypeName | Exprize | Title }}{Value: v, Op: op, Column: \"{{ .Name }}\"}\n\tq.where = append(q.where, where)\n\treturn q\n}\n\n{{ end }}\n"
var _Assets422c8ed1da204eae59c5a68d10bc81d2a9baa169 = "{{ define \"Delete\" }}\n{{- $camelName := .Name | toCamel | Untitle -}}\n{{- $constructor := printf \"New%sSQL\" (.Name | toCamel | Title) -}}\ntype {{ $camelName }}DeleteSQL struct {\n\t{{ $camelName }}SQL\n}\n\nfunc (q {{ $camelName }}SQL) Delete() {{ $camelName }}DeleteSQL {\n\treturn {{ $camelName }}DeleteSQL{\n\t\tq,\n\t}\n}\n\n{{ range .Columns }}{{ template \"DeleteColumn\" . }}{{ end }}\nfunc (q {{ $camelName }}DeleteSQL) ToSql() (string, []interface{}, error) {\n\twheres, vs, err := q.where.ToSql()\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\n\tquery := \"DELETE FROM {{ .Name }}\"\n\tif wheres != \"\" {\n\t\tquery += \" WHERE\" + wheres\n\t}\n\n\treturn query + \";\", vs, nil\n}\n\nfunc ( q {{ $camelName }}DeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn db.Exec(query, args...)\n}\n\n{{- if .HasPk }}\nfunc (s {{ .StructName }}) Delete(db sqlla.DB) (sql.Result, error) {\n\tquery, args, err := {{ $constructor }}().Delete().{{ .PkColumn.Name | toCamel | Title }}(s.{{ .PkColumn.FieldName }}).ToSql()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn db.Exec(query, args...)\n}\n{{- end }}\n{{ end }}\n"
var _Assetsfffa7c2a7f04a3a56a0a82e844d74ba9351ebec1 = "{{ define \"Insert\" }}\n{{- $camelName := .Name | toCamel | Untitle -}}\n{{- $constructor := printf \"New%sSQL\" (.Name | toCamel) -}}\ntype {{ $camelName }}InsertSQL struct {\n\t{{ $camelName }}SQL\n\tsetMap\tsqlla.SetMap\n\tColumns []string\n}\n\nfunc (q {{ $camelName }}SQL) Insert() {{ $camelName }}InsertSQL {\n\treturn {{ $camelName }}InsertSQL{\n\t\t{{ $camelName }}SQL: q,\n\t\tsetMap: sqlla.SetMap{},\n\t}\n}\n\n{{ range .Columns }}{{ template \"InsertColumn\" . }}{{ end }}\nfunc (q {{ $camelName }}InsertSQL) ToSql() (string, []interface{}, error) {\n\tvar err error\n\tvar s interface{} = {{ .StructName }}{}\n\tif t, ok := s.({{ $camelName }}DefaultInsertHooker); ok {\n\t\tq, err = t.DefaultInsertHook(q)\n\t\tif err != nil {\n\t\t\treturn \"\", []interface{}{}, err\n\t\t}\n\t}\n\tqs, vs, err := q.setMap.ToInsertSql()\n\tif err != nil {\n\t\treturn \"\", []interface{}{}, err\n\t}\n\n\tquery := \"INSERT INTO {{ .Name }} \" + qs\n\n\treturn query + \";\", vs, nil\n}\n\n{{ if .HasPk -}}\nfunc (q {{ $camelName }}InsertSQL) Exec(db sqlla.DB) ({{ .StructName }}, error) {\n{{- else -}}\nfunc (q {{ $camelName }}InsertSQL) Exec(db sqlla.DB) (sql.Result, error) {\n{{- end }}\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\t{{ if .HasPk -}}\n\t\treturn {{ .StructName }}{}, err\n\t\t{{- else }}\n\t\treturn nil, err\n\t\t{{- end }}\n\t}\n\tresult, err := db.Exec(query, args...)\n\t{{ if .HasPk -}}\n\tif err != nil {\n\t\treturn {{ .StructName }}{}, err\n\t}\n\tid, err := result.LastInsertId()\n\tif err != nil {\n\t\treturn {{ .StructName }}{}, err\n\t}\n\treturn {{ $constructor }}().Select().PkColumn(id).Single(db)\n\t{{- else -}}\n\treturn result, err\n\t{{- end }}\n}\n\ntype {{ $camelName }}DefaultInsertHooker interface {\n\tDefaultInsertHook({{ $camelName }}InsertSQL) ({{ $camelName }}InsertSQL, error)\n}\n{{ end }}\n"
var _Assets34a31430cf90520b4fc16b1999d6dafc4299a260 = "{{ define \"Select\" }}\n{{- $camelName := .Name | toCamel | Untitle -}}\n{{- $constructor := printf \"New%sSQL\" (.Name | toCamel | Title) -}}\nvar {{ $camelName }}AllColumns = []string{\n\t{{ range .Columns }}\"{{ .Name }}\",{{ end }}\n}\n\ntype {{ $camelName }}SelectSQL struct {\n\t{{ $camelName }}SQL\n\tColumns     []string\n\torder       string\n\tlimit       *uint64\n\tisForUpdate bool\n}\n\nfunc (q {{ $camelName }}SQL) Select() {{ $camelName }}SelectSQL {\n\treturn {{ $camelName }}SelectSQL{\n\t\tq,\n\t\t{{ $camelName }}AllColumns,\n\t\t\"\",\n\t\tnil,\n\t\tfalse,\n\t}\n}\n\nfunc (q {{ $camelName }}SelectSQL) Or(qs ...{{ $camelName }}SelectSQL) {{ $camelName }}SelectSQL {\n\tws := make([]sqlla.Where, 0, len(qs))\n\tfor _, q := range qs {\n\t\tws = append(ws, q.where)\n\t}\n\tq.where = append(q.where, sqlla.ExprOr(ws))\n\treturn q\n}\n\nfunc (q {{ $camelName }}SelectSQL) Limit(l uint64) {{ $camelName }}SelectSQL {\n\tq.limit = &l\n\treturn q\n}\n\nfunc (q {{ $camelName}}SelectSQL) ForUpdate() {{ $camelName }}SelectSQL {\n\tq.isForUpdate = true\n\treturn q\n}\n\n{{ range .Columns }}{{ template \"SelectColumn\" . }}{{ end }}\nfunc (q {{ $camelName }}SelectSQL) ToSql() (string, []interface{}, error) {\n\tcolumns := strings.Join(q.Columns, \", \")\n\twheres, vs, err := q.where.ToSql()\n\tif err != nil {\n\t\treturn \"\", nil, err\n\t}\n\n\tquery := \"SELECT \" + columns + \" FROM {{ .Name }}\"\n\tif wheres != \"\" {\n\t\tquery += \" WHERE\" + wheres\n\t}\n\tquery += q.order\n\tif q.limit != nil {\n\t\tquery += \" LIMIT \" + strconv.FormatUint(*q.limit, 10)\n\t}\n\n\tif q.isForUpdate {\n\t\tquery += \" FOR UPDATE\"\n\t}\n\n\treturn query + \";\", vs, nil\n}\n\n{{ if .HasPk -}}\nfunc (s {{ .StructName }}) Select() ({{ $camelName }}SelectSQL) {\n\treturn {{ $constructor }}().Select().{{ .PkColumn.Name | toCamel | Title }}(s.{{ .PkColumn.FieldName }})\n}\n{{ end -}}\n\nfunc (q {{ $camelName }}SelectSQL) Single(db sqlla.DB) ({{ .StructName }}, error) {\n\tq.Columns = {{ $camelName }}AllColumns\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\treturn {{ .StructName }}{}, err\n\t}\n\n\trow := db.QueryRow(query, args...)\n\treturn q.Scan(row)\n}\n\nfunc (q {{ $camelName }}SelectSQL) All(db sqlla.DB) ([]{{ .StructName }}, error) {\n\trs := make([]{{ .StructName }}, 0, 10)\n\tq.Columns = {{ $camelName }}AllColumns\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\trows, err := db.Query(query, args...)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tdefer rows.Close()\n\tfor rows.Next() {\n\t\tr, err := q.Scan(rows)\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\trs = append(rs, r)\n\t}\n\treturn rs, nil\n}\n\nfunc (q {{ $camelName }}SelectSQL) Scan(s sqlla.Scanner) ({{ .StructName }}, error) {\n\tvar row {{ .StructName }}\n\terr := s.Scan(\n\t\t{{ range .Columns }}&row.{{ (index .Names 0).String }},\n\t\t{{ end }}\n\t)\n\treturn row, err\n}\n{{ end }}\n"
var _Assets66097d0bb9f0a76f337837b48f15dac752c422c9 = "{{ define \"Update\" }}{{ $camelName := .Name | toCamel | Untitle }}\n{{- $constructor := printf \"New%sSQL\" (.Name | toCamel | Title) -}}\ntype {{ $camelName }}UpdateSQL struct {\n\t{{ $camelName }}SQL\n\tsetMap\tsqlla.SetMap\n\tColumns []string\n}\n\nfunc (q {{ $camelName }}SQL) Update() {{ $camelName }}UpdateSQL {\n\treturn {{ $camelName }}UpdateSQL{\n\t\t{{ $camelName }}SQL: q,\n\t\tsetMap: sqlla.SetMap{},\n\t}\n}\n\n{{ range .Columns }}{{ template \"UpdateColumn\" . }}{{ end }}\nfunc (q {{ $camelName }}UpdateSQL) ToSql() (string, []interface{}, error) {\n\tvar err error\n\tvar s interface{} = {{ .StructName }}{}\n\tif t, ok := s.({{ $camelName }}DefaultUpdateHooker); ok {\n\t\tq, err = t.DefaultUpdateHook(q)\n\t\tif err != nil {\n\t\t\treturn \"\", []interface{}{}, err\n\t\t}\n\t}\n\tsetColumns, svs, err := q.setMap.ToUpdateSql()\n\tif err != nil {\n\t\treturn \"\", []interface{}{}, err\n\t}\n\twheres, wvs, err := q.where.ToSql()\n\tif err != nil {\n\t\treturn \"\", []interface{}{}, err\n\t}\n\n\tquery := \"UPDATE {{ .Name }} SET\" + setColumns\n\tif wheres != \"\" {\n\t\tquery += \" WHERE\" + wheres\n\t}\n\n\treturn query + \";\", append(svs, wvs...), nil\n}\n\n{{- if .HasPk }}\nfunc (s {{ .StructName }}) Update() {{ $camelName }}UpdateSQL {\n\treturn {{ $constructor }}().Update().Where{{ .PkColumn.Name | toCamel | Title }}(s.{{ .PkColumn.FieldName }})\n}\n\nfunc (q {{ $camelName }}UpdateSQL) Exec(db sqlla.DB) ([]{{ .StructName }}, error) {\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\t_, err = db.Exec(query, args...)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\tqq := q.{{ $camelName }}SQL\n\n\treturn qq.Select().All(db)\n}\n{{- else }}\nfunc (q {{ $camelName }}UpdateSQL) Exec(db sqlla.DB) (sql.Result, error) {\n\tquery, args, err := q.ToSql()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn db.Exec(query, args...)\n}\n{{- end }}\n\ntype {{ $camelName }}DefaultUpdateHooker interface {\n\tDefaultUpdateHook({{ $camelName }}UpdateSQL) ({{ $camelName }}UpdateSQL, error)\n}\n{{ end }}\n"
var _Assets6137de71599adeee9ccdf90c810e11de5e0f148c = "{{ define \"DeleteColumn\" }}{{ $smallTableName := .TableName | toCamel | Untitle }}\nfunc (q {{ $smallTableName }}DeleteSQL) {{ .Name | toCamel | Title }}(v {{ .TypeName }}, exprs ...sqlla.Operator) {{ $smallTableName }}DeleteSQL {\n\tvar op sqlla.Operator\n\tif len(exprs) == 0 {\n\t\top = sqlla.OpEqual\n\t} else {\n\t\top = exprs[0]\n\t}\n\n\twhere := sqlla.Expr{{ .TypeName | Exprize | Title }}{Value: v, Op: op, Column: \"{{ .Name }}\"}\n\tq.where = append(q.where, where)\n\treturn q\n}\n\n\nfunc (q {{ $smallTableName }}DeleteSQL) {{ .Name | toCamel | Title }}In(vs ...{{ .TypeName }}) {{ $smallTableName }}DeleteSQL {\n\twhere := sqlla.ExprMulti{{ .TypeName | Exprize | Title }}{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: \"{{ .Name }}\"}\n\tq.where = append(q.where, where)\n\treturn q\n}\n{{ end }}\n"
var _Assetsda0090768d1a813fc3dcf7333f348dd837a5f320 = "{{ define \"InsertColumn\" }}{{ $smallTableName := .TableName | toCamel | Untitle }}\nfunc (q {{ $smallTableName }}InsertSQL) Value{{ .Name | toCamel | Title }}(v {{ .TypeName }}) {{ $smallTableName }}InsertSQL {\n\tq.setMap[\"{{ .Name }}\"] = v\n\treturn q\n}\n\n{{ end }}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"template"}, "/template": []string{"delete.tmpl", "delete_column.tmpl", "insert.tmpl", "insert_column.tmpl", "select.tmpl", "select_column.tmpl", "table.tmpl", "update.tmpl", "update_column.tmpl"}}, map[string]*assets.File{
	"/template/insert.tmpl": &assets.File{
		Path:     "/template/insert.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assetsfffa7c2a7f04a3a56a0a82e844d74ba9351ebec1),
	}, "/template/select_column.tmpl": &assets.File{
		Path:     "/template/select_column.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assets88a39ad92bb2c8db797af1c5e70ac4d41dd78298),
	}, "/template/table.tmpl": &assets.File{
		Path:     "/template/table.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1477359550, 1477359550000000000),
		Data:     []byte(_Assets6b103fa3a9e63224b4df28a8ed3efae32053f5bc),
	}, "/template/update_column.tmpl": &assets.File{
		Path:     "/template/update_column.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assets83c5f0ba25bc78b44a12ea398292f04e3a78d0a0),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1518058157, 1518058157000000000),
		Data:     nil,
	}, "/template/delete.tmpl": &assets.File{
		Path:     "/template/delete.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assets422c8ed1da204eae59c5a68d10bc81d2a9baa169),
	}, "/template/insert_column.tmpl": &assets.File{
		Path:     "/template/insert_column.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assetsda0090768d1a813fc3dcf7333f348dd837a5f320),
	}, "/template/select.tmpl": &assets.File{
		Path:     "/template/select.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1517997636, 1517997636000000000),
		Data:     []byte(_Assets34a31430cf90520b4fc16b1999d6dafc4299a260),
	}, "/template/update.tmpl": &assets.File{
		Path:     "/template/update.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assets66097d0bb9f0a76f337837b48f15dac752c422c9),
	}, "/template": &assets.File{
		Path:     "/template",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1517997636, 1517997636000000000),
		Data:     nil,
	}, "/template/delete_column.tmpl": &assets.File{
		Path:     "/template/delete_column.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1474851811, 1474851811000000000),
		Data:     []byte(_Assets6137de71599adeee9ccdf90c810e11de5e0f148c),
	}}, "")
