package sqlla

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
