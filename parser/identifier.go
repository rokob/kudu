package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// IdentifierParslet - the parslet for handling identifiers
type IdentifierParslet struct{}

func (p *IdentifierParslet) parse(parser *Parser, token token.Token) ast.Expression {
	return ast.IdentifierExpression{Identifier: token.Literal}
}

func (p *IdentifierParslet) String() string {
	return "IdentifierParslet"
}
