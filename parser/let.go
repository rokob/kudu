package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// LetParslet - the parslet for handling variable bindings
type LetParslet struct{}

func (p *LetParslet) parse(parser *Parser, inputToken token.Token) ast.Expression {
	identifierToken := parser.consumeExpected(token.IDENT)
	if identifierToken.Type == token.ILLEGAL {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("Expected to see an identifier on the left-hand-side of a let statment, got something else")
		}
	}
	identifier := (&IdentifierParslet{}).parse(parser, identifierToken)
	if t := parser.consumeExpected(token.ASSIGN); t.Type == token.ILLEGAL {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("Expected to see an = in a let expression, but saw something else")
		}
	}
	binding := parser.parseExpression()
	if _, ok := binding.(ast.IllegalExpression); ok {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("The binding expression in a let statement is illegal")
		}
	}
	return ast.LetExpression{Type: ast.LetExpressionType, Identifier: identifier, Binding: binding}
}
