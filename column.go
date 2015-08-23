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

type SetMap map[string]interface{}

func (sm SetMap) ToSql() (string, []interface{}, error) {
	var setColumns string
	vs := []interface{}{}
	columnCount := 0
	for k, v := range sm {
		if columnCount != 0 {
			setColumns += ","
		}
		setColumns += " " + k + " = ?"
		vs = append(vs, v)
	}

	return setColumns, vs, nil
}
