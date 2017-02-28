package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// PrefixOperatorParslet - the parselt for handling prefix operators
type PrefixOperatorParslet struct{}

// PrefixExpression - an expression representing a prefix operator
type PrefixExpression struct {
	Type  token.Type
	Right Expression
}

func (e PrefixExpression) String() string {
	return fmt.Sprintf("PREFIX(%s, %s)", e.Type, e.Right.String())
}

func (p *PrefixOperatorParslet) parse(parser *Parser, token token.Token) Expression {
	right := parser.parseExpression()
	if _, ok := right.(IllegalExpression); ok {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("The right side of a prefix expression is illegal")
		}
	}
	return PrefixExpression{Type: token.Type, Right: right}
}

func (p *PrefixOperatorParslet) String() string {
	return "PrefixOperatorParslet"
}
