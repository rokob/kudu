package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// PrefixOperatorParslet - the parselt for handling prefix operators
type PrefixOperatorParslet struct{}

func (p *PrefixOperatorParslet) parse(parser *Parser, token token.Token) ast.Expression {
	right := parser.parseExpression()
	if _, ok := right.(ast.IllegalExpression); ok {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("The right side of a prefix expression is illegal")
		}
	}
	return ast.PrefixExpression{Type: token.Type, Right: right}
}

func (p *PrefixOperatorParslet) String() string {
	return "PrefixOperatorParslet"
}
