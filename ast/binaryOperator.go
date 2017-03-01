package ast

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// OperatorExpressionArguments is the type for arguments to a binary operator
type OperatorExpressionArguments struct {
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

// OperatorExpression - an expression for a binary operator
type OperatorExpression struct {
	Type      token.Type                  `json:"binary operator"`
	Arguments OperatorExpressionArguments `json:"args"`
}

func (e OperatorExpression) String() string {
	return fmt.Sprintf("BINARY(%s, %s, %s)", e.Type, e.Arguments.Left.String(), e.Arguments.Right.String())
}

// Visit this AST Node
func (e OperatorExpression) Visit(env Environment) Value {
	leftValue := e.Arguments.Left.Visit(env)
	rightValue := e.Arguments.Right.Visit(env)

	if leftValue.Type != "int" && rightValue.Type != "int" {
		return ErrorValue("I only understand integer math right now")
	}
	left := leftValue.Data.(int64)
	right := rightValue.Data.(int64)
	switch e.Type {
	case token.PLUS:
		return Value{Type: "int", Data: left + right}
	case token.MINUS:
		return Value{Type: "int", Data: left - right}
	default:
		return ErrorValue(fmt.Sprintf("Unknown operator type: %s", e.Type))
	}
}
