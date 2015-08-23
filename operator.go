package sqlla

var (
	OpEqual        Operator = "="
	OpGreater      Operator = ">"
	OpGreaterEqual Operator = ">="
	OpLess         Operator = "<"
	OpLessEqual    Operator = "<="
	OpNot          Operator = "<>"
	OpIs           Operator = "IS"
	OpIsNull       Operator = "IS NULL"
)

type Operator string

func (op Operator) ToSql() (string, error) {
	return string(op), nil
}
