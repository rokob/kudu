package ast

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// PrefixExpression - an expression representing a prefix operator
type PrefixExpression struct {
	Type  token.Type `json:"prefix operator"`
	Right Expression `json:"operand"`
}

func (e PrefixExpression) String() string {
	return fmt.Sprintf("PREFIX(%s, %s)", e.Type, e.Right.String())
}

// Visit this AST Node
func (e PrefixExpression) Visit(env Environment) Value {
	return NoneValue()
}
