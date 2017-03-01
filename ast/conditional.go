package ast

import (
	"fmt"
)

// ConditionalExpression - an expression representing an if/else statement
type ConditionalExpression struct {
	IfCondition   Expression             `json:"if"`
	Block         []Expression           `json:"then"`
	ElseCondition *ConditionalExpression `json:"else"`
}

func (e ConditionalExpression) String() string {
	if e.IfCondition != nil {
		if e.ElseCondition == nil {
			return fmt.Sprintf("IF(%s, THEN(%s))", e.IfCondition, expressionListJoin(e.Block, ";"))
		}
		return fmt.Sprintf("IF(%s, THEN(%s), ELSE(%s))", e.IfCondition, expressionListJoin(e.Block, ";"), *e.ElseCondition)
	}
	return expressionListJoin(e.Block, ";")
}

// Visit this AST Node
func (e ConditionalExpression) Visit(env Environment) Value {
	return NoneValue()
}
