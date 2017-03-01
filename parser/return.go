package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// ReturnParslet - the parselt for handling prefix operators
type ReturnParslet struct{}

func (p *ReturnParslet) parse(parser *Parser, inputToken token.Token) ast.Expression {
	expression := parser.parseExpression()
	if _, ok := expression.(ast.IllegalExpression); ok {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("An expression being returnd is illegal")
		}
	}
	return ast.ReturnExpression{Value: expression}
}

func (p *ReturnParslet) String() string {
	return "ReturnParslet"
}
