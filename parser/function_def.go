package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// FunctionDefinitionParslet - the parslet for handling function definitions
type FunctionDefinitionParslet struct{}

func (p *FunctionDefinitionParslet) parse(parser *Parser, inputToken token.Token) ast.Expression {
	arguments := make([]ast.Expression, 0)
	block := make([]ast.Expression, 0)
	if !parser.match(token.LBRACE) {
		if t := parser.consumeExpected(token.LPAREN); t.Type == token.ILLEGAL {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("Expected to see a ( in a function defintiion, got something else")
			}
		}
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
				panic("Expected to see a ) in a function defintiion, got something else")
			}
		}
		if t := parser.consumeExpected(token.LBRACE); t.Type == token.ILLEGAL {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("Expected to see a { in a function defintiion, got something else")
			}
		}
	}
	for !parser.match(token.RBRACE) {
		b := parser.parseExpression()
		if _, ok := b.(ast.IllegalExpression); ok {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("An expression in a function definition block is illegal")
			}
		}
		block = append(block, b)
	}
	return ast.FunctionDefinitionExpression{Type: ast.FunctionDefinitionExpressionType, Arguments: arguments, Block: block}
}
