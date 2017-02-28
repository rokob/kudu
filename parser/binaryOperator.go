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

type operatorExpressionArguments struct {
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

// OperatorExpression - an expression for a binary operator
type OperatorExpression struct {
	Type      token.Type                  `json:"binary operator"`
	Arguments operatorExpressionArguments `json:"args"`
}

func (e OperatorExpression) String() string {
	return fmt.Sprintf("BINARY(%s, %s, %s)", e.Type, e.Arguments.Left.String(), e.Arguments.Right.String())
}

func (p *BinaryOperatorParselet) parse(parser *Parser, left Expression, token token.Token) Expression {
	precedence := p.precedence
	if p.isRight {
		precedence--
	}
	if p.leftCondition != nil {
		if ok, msg := p.leftCondition(left); !ok {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic(msg)
			}
		}
	}
	right := parser.parseExpression(precedence)
	if _, ok := right.(IllegalExpression); ok {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic(fmt.Sprintf("The condition on the right side of the binary operator %s is illegal", token.Literal))
		}
	}
	if p.rightCondition != nil {
		if ok, msg := p.rightCondition(right); !ok {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic(msg)
			}
		}
	}
	return OperatorExpression{
		Type: token.Type,
		Arguments: operatorExpressionArguments{
			Left:  left,
			Right: right,
		},
	}
}

func (p *BinaryOperatorParselet) getPrecedence() Precedence {
	return p.precedence
}

func (p *BinaryOperatorParselet) String() string {
	return fmt.Sprintf("BinaryOperatorParselet(precedence: %d, isRight: %t)", p.precedence, p.isRight)
}
