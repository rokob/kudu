package parser

import (
	"github.com/rokob/kudu/token"
)

// GroupParslet - the parselt for handling prefix operators
type GroupParslet struct{}

func (p *GroupParslet) parse(parser *Parser, inputToken token.Token) Expression {
	expression := parser.parseExpression()
	parser.consumeExpected(token.RPAREN)
	return expression
}

func (p *GroupParslet) String() string {
	return "GroupParslet"
}
