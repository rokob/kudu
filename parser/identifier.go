package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// IdentifierParslet - the parslet for handling identifiers
type IdentifierParslet struct{}

// IdentifierExpression - an expression representing an identifier
type IdentifierExpression struct {
	Identifier string
}

func (e IdentifierExpression) String() string {
	return fmt.Sprintf("IDENT(%s)", e.Identifier)
}

func (p *IdentifierParslet) parse(parser *Parser, token token.Token) Expression {
	return IdentifierExpression{Identifier: token.Literal}
}

func (p *IdentifierParslet) String() string {
	return "IdentifierParslet"
}
