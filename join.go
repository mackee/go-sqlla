package sqlla

import (
	"go/ast"
	"go/types"
)

type toJoinDefinitionOpts struct {
	defPkg            *types.Package
	annotationComment string
	gd                *ast.GenDecl
	ti                *types.Info
}

// JoinDefinition is definition of join statement
type JoinDefinition struct {
	Name       string
	Columns    map[string][]string
	JoinTables []JoinTable
}

// JoinTable is definition of table information for join
type JoinTable struct {
	TableName string
	AltName   string
	Condition []EvaluatedJoinDefinition
	Table     Table
}

// JoinKind is kind of JOIN
type JoinKind int

const (
	// JoinTypeInner is INNER JOIN
	JoinTypeInner JoinKind = iota + 1
	// JoinTypeLeftOuter is LEFT OUTER JOIN
	JoinTypeLeftOuter
	// JoinTypeRightOuter is RIGHT OUTER JOIN
	JoinTypeRightOuter
)

// EvaluatedJoinDefinition is definition of join condition after evaluated
type EvaluatedJoinDefinition struct {
	Kind    JoinKind
	Column1 string
	Column2 string
}

func toJoinDefinition(opts toJoinDefinitionOpts) (*JoinDefinition, error) {
	return nil, nil
}

// InnerJoiner is implements definition for INNER JOIN
type InnerJoiner interface {
	_innerJoin(methods JoinMethods) []JoinCondition
}

// LeftOuterJoiner is implements definition for LEFT OUTER JOIN
type LeftOuterJoiner interface {
	_leftOuterJoin(methods JoinMethods) []JoinCondition
}

// RightOuterJoiner is implements definition for RIGHT OUTER JOIN
type RightOuterJoiner interface {
	_rightOuterJoin(methods JoinMethods) []JoinCondition
}

// JoinMethods provides DSL for JOIN at the schema
type JoinMethods interface {
	On(bool) JoinCondition
}

// JoinCondition provides DSL for JOIN at the schema
type JoinCondition interface{}
