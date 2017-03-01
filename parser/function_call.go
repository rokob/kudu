package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// FunctionCallParslet - the parslet for handling function calls
type FunctionCallParslet struct{}

func (p *FunctionCallParslet) parse(parser *Parser, left ast.Expression, inputToken token.Token) ast.Expression {
	arguments := make([]ast.Expression, 0)
	if !parser.match(token.RPAREN) {
		firstArgument := parser.parseExpression()
		if _, ok := firstArgument.(ast.IllegalExpression); ok {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("The expression for the first argument in a function definition is illegal")
			}
		}
		arguments = append(arguments, firstArgument)
		for parser.match(token.COMMA) {
			a := parser.parseExpression()
			if _, ok := a.(ast.IllegalExpression); ok {
				if parser.mode == ReplMode {
					return ast.IllegalExpression{}
				} else if parser.mode == CompilerMode {
					panic("The expression for an argument in a function definition is illegal")
				}
			}
			arguments = append(arguments, a)
		}
	}
	if t := parser.consumeExpected(token.RPAREN); t.Type == token.ILLEGAL {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("Expected to see a ) in a function call, got something else")
		}
	}

	return ast.FunctionCallExpression{Type: ast.FunctionCallExpressionType, Function: left, Arguments: arguments}
}

func (p *FunctionCallParslet) getPrecedence() Precedence {
	return CALL
}
