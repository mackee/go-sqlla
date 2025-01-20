package sqlla

import (
	"fmt"
	"go/ast"
	"sort"
	"strconv"
	"strings"
)

type Columns []Column

type Column struct {
	Field        *ast.Field
	Name         string
	MethodName   string
	typeName     string
	PkgName      string
	baseTypeName string
	altTypeName  string
	TableName    string
	IsPk         bool
	isNullT      bool
}

func (c Column) HasUnderlyingType() bool {
	return c.baseTypeName != c.typeName
}

func (c Column) TypeName() string {
	tn := c.typeName
	return tn
}

func (c Column) BaseTypeName() string {
	return c.baseTypeName
}

func (c Column) AltTypeName() string {
	if c.altTypeName == "" {
		return ""
	}
	return c.altTypeName
}

func (c Column) IsNullT() bool {
	return c.isNullT
}

func (c Column) nullTypeSuffix() string {
	nv := strings.TrimPrefix(c.baseTypeName, "sql.Null")
	nv = strings.TrimPrefix(nv, "mysql.Null")
	if nv == c.baseTypeName {
		return ""
	}
	return nv
}

func (c Column) nullBaseType(t string) string {
	if t == "" {
		return ""
	}
	if t == "Time" {
		return "time.Time"
	}

	return strings.ToLower(t)
}

func (c Column) ExprValue() string {
	if nv := c.nullBaseType(c.nullTypeSuffix()); nv != "" {
		return "sqlla.ExprNull[" + nv + "]"
	}
	if c.isNullT {
		return "sqlla.ExprNull[" + c.baseTypeName + "]"
	}
	return "sqlla.ExprValue[" + c.baseTypeName + "]"
}

func (c Column) ExprMultiValue() string {
	return "sqlla.ExprMultiValue[" + c.baseTypeName + "]"
}

func (c Column) ExprValueIdentifier() string {
	if nt := c.nullTypeSuffix(); nt != "" {
		return "sql.Null[" + c.nullBaseType(nt) + "]{Valid: v.Valid, V: v." + nt + "}"
	}
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

type Where []Expr

func (wh Where) ToSql() (string, []any, error) {
	if len(wh) == 0 {
		return "", nil, nil
	}
	wheres := " "
	vs := []any{}
	for i, w := range wh {
		s, v, err := w.ToSql()
		if err != nil {
			return "", nil, err
		}
		vs = append(vs, v...)

		if i == 0 {
			wheres += s
			continue
		}
		wheres += " AND " + s
	}

	return wheres, vs, nil
}

func (wh Where) ToSqlPg(offset int) (string, int, []any, error) {
	if len(wh) == 0 {
		return "", offset, nil, nil
	}
	wheres := " "
	vs := []any{}
	for i, w := range wh {
		s, n, v, err := w.ToSqlPg(offset)
		if err != nil {
			return "", offset, nil, err
		}
		offset = n
		vs = append(vs, v...)
		if i == 0 {
			wheres += s
			continue
		}
		wheres += " AND " + s
	}

	return wheres, offset, vs, nil
}

type SetMapRawValue string

type SetMap map[string]any

func (sm SetMap) NewIterator() *SetMapIterator {
	keys := make(sort.StringSlice, 0, len(sm))
	for k := range sm {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return &SetMapIterator{
		sm:     sm,
		keys:   keys,
		cursor: -1,
	}
}

type SetMapIterator struct {
	sm     SetMap
	cursor int
	keys   []string
}

func (s *SetMapIterator) Iterate() bool {
	s.cursor++
	return len(s.keys)-1 >= s.cursor
}

func (s *SetMapIterator) Key() string {
	return s.keys[s.cursor]
}

func (s *SetMapIterator) Value() any {
	return s.sm[s.keys[s.cursor]]
}

// ToUpdateSqlPg generates to set values SQL expressions with placeholders for MySQL/SQLite.
func (sm SetMap) ToUpdateSql() (string, []any, error) {
	var setColumns string
	vs := []any{}
	columnCount := 0
	iter := sm.NewIterator()
	for iter.Iterate() {
		k, v := iter.Key(), iter.Value()
		if columnCount != 0 {
			setColumns += ","
		}
		if rv, ok := v.(SetMapRawValue); ok {
			setColumns += " " + k + " = " + string(rv)
		} else {
			setColumns += " " + k + " = ?"
			vs = append(vs, v)
		}
		columnCount++
	}

	return setColumns, vs, nil
}

// ToUpdateSqlPg generates to set values SQL expressions with numbered placeholders for PostgreSQL.
func (sm SetMap) ToUpdateSqlPg(offset int) (string, int, []any, error) {
	var setColumns string
	vs := []any{}
	columnCount := 0
	placeholderNum := offset
	iter := sm.NewIterator()
	for iter.Iterate() {
		k, v := iter.Key(), iter.Value()
		if columnCount != 0 {
			setColumns += ","
		}
		if rv, ok := v.(SetMapRawValue); ok {
			setColumns += " " + k + " = " + string(rv)
		} else {
			placeholderNum++
			setColumns += " " + k + " = $" + strconv.Itoa(placeholderNum)
			vs = append(vs, v)
		}
		columnCount++
	}

	return setColumns, placeholderNum, vs, nil
}

// ToInsertColumnsAndValues generates to insert columns and values SQL expressions with placeholders.
func (sm SetMap) ToInsertColumnsAndValues() (string, string, []any) {
	qs, ps := "(", "("
	vs := []any{}
	columnCount := 0
	iter := sm.NewIterator()
	for iter.Iterate() {
		k, v := iter.Key(), iter.Value()
		if columnCount != 0 {
			qs += ","
			ps += ","
		}
		qs += k
		ps += "?"
		vs = append(vs, v)
		columnCount++
	}
	qs += ")"
	ps += ")"
	return qs, ps, vs
}

// ToInsertColumnAndValuesPg generates to insert columns and values SQL expressions with numbered placeholders for PostgreSQL.
func (sm SetMap) ToInsertColumnsAndValuesPg(offset int) (string, string, int, []any) {
	qs, ps := "(", "("
	vs := []any{}
	columnCount := 0
	placeholderNum := offset
	iter := sm.NewIterator()
	for iter.Iterate() {
		k, v := iter.Key(), iter.Value()
		if columnCount != 0 {
			qs += ","
			ps += ","
		}
		qs += k
		placeholderNum++
		columnCount++
		ps += "$" + strconv.Itoa(placeholderNum)
		vs = append(vs, v)
	}
	qs += ")"
	ps += ")"
	return qs, ps, placeholderNum, vs
}

// ToInsertSql generates to insert SQL expressions with placeholders.
func (sm SetMap) ToInsertSql() (string, []any, error) {
	qs, ps, vs := sm.ToInsertColumnsAndValues()
	return qs + " VALUES " + ps, vs, nil
}

// ToInsertSqlPg generates to insert SQL expressions with numbered placeholders for PostgreSQL.
func (sm SetMap) ToInsertSqlPg(offset int) (string, int, []any, error) {
	qs, ps, placeholderNum, vs := sm.ToInsertColumnsAndValuesPg(offset)
	return qs + " VALUES " + ps, placeholderNum, vs, nil
}

type SetMaps []SetMap

// ToInsertSql generates to insert SQL expressions with placeholders.
func (s SetMaps) ToInsertSql() (string, []any, error) {
	if len(s) == 0 {
		return "", nil, fmt.Errorf("sqlla: SetMaps is empty")
	}

	first := s[0]
	columns, values, vs := first.ToInsertColumnsAndValues()
	var b strings.Builder
	if _, err := b.WriteString(values); err != nil {
		return "", nil, err
	}
	for i, _s := range s[1:] {
		_columns, _values, _vs := _s.ToInsertColumnsAndValues()
		if columns != _columns {
			return "", nil, fmt.Errorf("sqlla: two SetMap are not match keys: [0]=%s, [%d]=%s", columns, i, _columns)
		}
		vs = append(vs, _vs...)
		if _, err := b.WriteString(","); err != nil {
			return "", nil, err
		}
		if _, err := b.WriteString(_values); err != nil {
			return "", nil, err
		}
	}
	return columns + " VALUES " + b.String(), vs, nil
}

// ToInsertSqlPg generates to insert SQL expressions with numbered placeholders for PostgreSQL.
func (s SetMaps) ToInsertSqlPg(offset int) (string, int, []any, error) {
	if len(s) == 0 {
		return "", 0, nil, fmt.Errorf("sqlla: SetMaps is empty")
	}

	first := s[0]
	columns, values, placeholderNum, vs := first.ToInsertColumnsAndValuesPg(offset)
	var b strings.Builder
	if _, err := b.WriteString(values); err != nil {
		return "", 0, nil, err
	}
	for i, _s := range s[1:] {
		_columns, _values, _placeholderNum, _vs := _s.ToInsertColumnsAndValuesPg(placeholderNum)
		if columns != _columns {
			return "", 0, nil, fmt.Errorf("sqlla: two SetMap are not match keys: [0]=%s, [%d]=%s", columns, i, _columns)
		}
		vs = append(vs, _vs...)
		if _, err := b.WriteString(","); err != nil {
			return "", 0, nil, err
		}
		if _, err := b.WriteString(_values); err != nil {
			return "", 0, nil, err
		}
		placeholderNum = _placeholderNum
	}
	return columns + " VALUES " + b.String(), placeholderNum, vs, nil
}
