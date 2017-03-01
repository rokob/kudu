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
