package ast

import (
	"fmt"
)

// ReturnExpression - an expression representing a return from a function
type ReturnExpression struct {
	Value Expression `json:"return"`
}

func (e ReturnExpression) String() string {
	return fmt.Sprintf("RETURN(%s)", e.Value)
}

