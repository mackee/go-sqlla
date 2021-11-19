package sqlla

type Where []Expr

func (wh Where) ToSql() (string, []interface{}, error) {
	if len(wh) == 0 {
		return "", nil, nil
	}
	wheres := " "
	vs := []interface{}{}
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

type SetMapRawValue string

type SetMap map[string]interface{}

func (sm SetMap) ToUpdateSql() (string, []interface{}, error) {
	var setColumns string
	vs := []interface{}{}
	columnCount := 0
	for k, v := range sm {
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

func (sm SetMap) ToInsertSql() (string, []interface{}, error) {
	qs, ps := "(", "("
	vs := []interface{}{}
	columnCount := 0
	for k, v := range sm {
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

	return qs + " VALUES" + ps, vs, nil
}
