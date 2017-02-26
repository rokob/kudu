package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// LetParslet - the parslet for handling variable bindings
type LetParslet struct{}

// LetExpression - an expression representing a let binding
type LetExpression struct {
	Identifier Expression
	Binding    Expression
}

func (e LetExpression) String() string {
	return fmt.Sprintf("LET(%s = %s)", e.Identifier, e.Binding)
}

func (p *LetParslet) parse(parser *Parser, inputToken token.Token) Expression {
	identifierToken := parser.consumeExpected(token.IDENT)
	identifier := (&IdentifierParslet{}).parse(parser, identifierToken)
	parser.consumeExpected(token.ASSIGN)
	binding := parser.parseExpression()
	return LetExpression{Identifier: identifier, Binding: binding}
}
