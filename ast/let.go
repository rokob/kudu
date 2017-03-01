package ast

import (
	"fmt"
)

// LetExpressionType is a string for debugging let definitions
const LetExpressionType string = "variable declaration"

// LetExpression - an expression representing a let binding
type LetExpression struct {
	Type       string     `json:"type"`
	Identifier Expression `json:"identifier"`
	Binding    Expression `json:"binding"`
}

func (e LetExpression) String() string {
	return fmt.Sprintf("LET(%s = %s)", e.Identifier, e.Binding)
}

// Visit this AST Node
func (e LetExpression) Visit(env Environment) Value {
	return NoneValue()
}
