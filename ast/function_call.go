package ast

import (
	"fmt"
)

// FunctionCallExpressionType is the type for a function call
const FunctionCallExpressionType string = "function call"

// FunctionCallExpression - an expression representing a function call
type FunctionCallExpression struct {
	Type      string       `json:"type"`
	Function  Expression   `json:"name"`
	Arguments []Expression `json:"args"`
}

func (e FunctionCallExpression) String() string {
	return fmt.Sprintf("CALL(%s;(%s))", e.Function, expressionListJoin(e.Arguments, ","))
}

// Visit this AST Node
func (e FunctionCallExpression) Visit(env Environment) Value {
	return NoneValue()
}
