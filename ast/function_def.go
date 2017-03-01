package ast

import (
	"fmt"
)

// FunctionDefinitionExpressionType is a string for debugging a function definition expression
const FunctionDefinitionExpressionType string = "function definition"

// FunctionDefinitionExpression - an expression representing a function definition
type FunctionDefinitionExpression struct {
	Type      string       `json:"type"`
	Arguments []Expression `json:"args"`
	Block     []Expression `json:"body"`
}

func (e FunctionDefinitionExpression) String() string {
	return fmt.Sprintf("FUN((%s);(%s))", expressionListJoin(e.Arguments, ","), expressionListJoin(e.Block, ";"))
}
