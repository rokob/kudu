package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// BinaryOperatorParselet - parslet for binary operators
type BinaryOperatorParselet struct {
	precedence     Precedence
	isRight        bool
	leftCondition  func(Expression) (bool, string)
	rightCondition func(Expression) (bool, string)
}

// OperatorExpression - an expression for a binary operator
type OperatorExpression struct {
	Left  Expression
	Type  token.Type
	Right Expression
}

func (e OperatorExpression) String() string {
	return fmt.Sprintf("BINARY(%s, %s, %s)", e.Type, e.Left.String(), e.Right.String())
}

func (p *BinaryOperatorParselet) parse(parser *Parser, left Expression, token token.Token) Expression {
	precedence := p.precedence
	if p.isRight {
		precedence--
	}
	if p.leftCondition != nil {
		if ok, msg := p.leftCondition(left); !ok {
			panic(msg)
		}
	}
	right := parser.parseExpression(precedence)
	if p.rightCondition != nil {
		if ok, msg := p.rightCondition(right); !ok {
			panic(msg)
		}
	}
	return OperatorExpression{Left: left, Type: token.Type, Right: right}
}

func (p *BinaryOperatorParselet) getPrecedence() Precedence {
	return p.precedence
}

func (p *BinaryOperatorParselet) String() string {
	return fmt.Sprintf("BinaryOperatorParselet(precedence: %d, isRight: %t)", p.precedence, p.isRight)
}
