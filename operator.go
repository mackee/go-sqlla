package sqlla

import (
	"strconv"
	"strings"
)

var (
	OpEqual                        OperatorBinary = "="  // Operator for equal. Column(value, sqlla.OpEqual) same as Column = value
	OpGreater                      OperatorBinary = ">"  // Operator for greater. Column(value, sqlla.OpGreater) same as Column > value
	OpGreaterEqual                 OperatorBinary = ">=" // Operator for greater equal. Column(value, sqlla.OpGreaterEqual) same as Column >= value
	OpLess                         OperatorBinary = "<"  // Operator for less. Column(value, sqlla.OpLess) same as Column < value
	OpLessEqual                    OperatorBinary = "<=" // Operator for less equal. Column(value, sqlla.OpLessEqual) same as Column <= value
	OpNot                          OperatorBinary = "<>" // Operator for not equal. Column(value, sqlla.OpNot) same as Column <> value
	OpIs                           OperatorBinary = "IS" // Operator for is. Column(value, sqlla.OpIs) same as Column IS value
	opIsNull                       OperatorBinary = "IS NULL"
	opIsNotNull                    OperatorBinary = "IS NOT NULL"
	OpLike                         OperatorBinary = "LIKE" // Operator for like. Column(value, sqlla.OpLike) same as Column LIKE value
	OpPgvectorL2                   OperatorBinary = "<->"  // Operator for pgvector l2 distance. Column(value, sqlla.OpPgvectorL2) same as Column <-> value
	OpPgvectorNegativeInnerProduct OperatorBinary = "<#>"  // Operator for pgvector negative inner product. Column(value, sqlla.OpPgvectorNegativeInnerProduct) same as Column <#> value
	OpPgvectorCosine               OperatorBinary = "<=>"  // Operator for pgvector cosine distance. Column(value, sqlla.OpPgvectorCosineDistance) same as Column <=> value
	OpPgvectorL1                   OperatorBinary = "<+>"  // Operator for pgvector l1 distance. Column(value, sqlla.OpPgvectorL1) same as Column <+> value
	OpPgvectorHamming              OperatorBinary = "<~>"  // Operator for pgvector hamming distance. Column(value, sqlla.OpPgvectorHamming) same as Column <~> value
	OpPgvectorJaccard              OperatorBinary = "<%>"  // Operator for pgvector jaccard distance. Column(value, sqlla.OpPgvectorJaccard) same as Column <%> value

)

type OperatorBinary string

func (op OperatorBinary) ToSql() (string, error) {
	return string(op), nil
}

type Operator interface {
	ToSql() (string, error)
}

type OperatorMulti interface {
	Operator
	ToSqlPg(offset int) (string, int, error)
}

type OperatorIn struct {
	num int
}

func (o *OperatorIn) String() string {
	return "IN(?" + strings.Repeat(",?", o.num-1) + ")"
}

func (o *OperatorIn) ToSql() (string, error) {
	return o.String(), nil
}

func (o *OperatorIn) ToSqlPg(offset int) (string, int, error) {
	b := &strings.Builder{}
	b.WriteString("IN(")
	for i := range o.num {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString("$" + strconv.Itoa(offset+i+1))
	}
	b.WriteString(")")
	return b.String(), offset + o.num, nil
}

func MakeInOperator(n int) *OperatorIn {
	return &OperatorIn{num: n}
}

type operatorAndValue struct {
	Op    OperatorBinary
	Value any
}

func NewOperatorAndValue(op OperatorBinary, value any) OperatorAndValues {
	return &operatorAndValue{Op: op, Value: value}
}

type OperatorAndValues interface {
	ToSql() (string, error)
	Values() []any
}

func (o *operatorAndValue) ToSql() (string, error) {
	return o.Op.ToSql()
}

func (o *operatorAndValue) Values() []any {
	return []any{o.Value}
}

type Operators []Operator

func (o Operators) ToSql() (string, error) {
	if len(o) == 0 {
		return OpEqual.ToSql()
	}
	return o[0].ToSql()
}

func (o Operators) Values() []any {
	vs := make([]any, 0, len(o))
	for _, op := range o {
		oav, ok := op.(OperatorAndValues)
		if !ok {
			continue
		}
		vs = append(vs, oav.Values()...)
	}
	return vs
}
