package sqlla

import "strconv"

const (
	Asc  OrderSimple = "ASC"
	Desc OrderSimple = "DESC"
)

type Order interface {
	WithColumn(columnName string) OrderWithColumn
}

type OrderWithColumn interface {
	OrderExpr() string
	OrderExprPg(offset int) (string, int)
	Values() []any
}

type OrderSimple string

type orderSimpleWithColumn struct {
	order      OrderSimple
	columnName string
}

func (o OrderSimple) WithColumn(columnName string) OrderWithColumn {
	return &orderSimpleWithColumn{order: o, columnName: columnName}
}

func (o *orderSimpleWithColumn) OrderExpr() string {
	return o.columnName + " " + string(o.order)
}

func (o *orderSimpleWithColumn) OrderExprPg(offset int) (string, int) {
	return o.columnName + " " + string(o.order), offset
}

func (o *orderSimpleWithColumn) Values() []any {
	return nil
}

type OrderWithOperator struct {
	op    OperatorBinary
	value any
	order OrderSimple
}

func NewOrderWithOperator(op OperatorBinary, value any, order OrderSimple) *OrderWithOperator {
	return &OrderWithOperator{op: op, value: value, order: order}
}

func (o *OrderWithOperator) WithColumn(columnName string) OrderWithColumn {
	return &orderWithOperatorAndColumn{owo: o, columnName: columnName}
}

type orderWithOperatorAndColumn struct {
	owo        *OrderWithOperator
	columnName string
}

func (o *orderWithOperatorAndColumn) OrderExpr() string {
	return o.columnName + " " + string(o.owo.op) + " ? " + string(o.owo.order)
}

func (o *orderWithOperatorAndColumn) OrderExprPg(offset int) (string, int) {
	return o.columnName + " " + string(o.owo.op) + " $" + strconv.Itoa(offset+1) + " " + string(o.owo.order), offset + 1
}

func (o *orderWithOperatorAndColumn) Values() []any {
	return []any{o.owo.value}
}
