package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// FunctionDefinitionParslet - the parslet for handling function definitions
type FunctionDefinitionParslet struct{}

const functionDefinitionExpressionType string = "function definition"

// FunctionDefinitionExpression - an expression representing a function definition
type FunctionDefinitionExpression struct {
	Type      string       `json:"type"`
	Arguments []Expression `json:"args"`
	Block     []Expression `json:"body"`
}

func (e FunctionDefinitionExpression) String() string {
	return fmt.Sprintf("FUN((%s);(%s))", expressionListJoin(e.Arguments, ","), expressionListJoin(e.Block, ";"))
}

func (p *FunctionDefinitionParslet) parse(parser *Parser, inputToken token.Token) Expression {
	arguments := make([]Expression, 0)
	block := make([]Expression, 0)
	if !parser.match(token.LBRACE) {
		if t := parser.consumeExpected(token.LPAREN); t.Type == token.ILLEGAL {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("Expected to see a ( in a function defintiion, got something else")
			}
		}
		if !parser.match(token.RPAREN) {
			firstArgument := parser.parseExpression()
			if _, ok := firstArgument.(IllegalExpression); ok {
				if parser.mode == ReplMode {
					return IllegalExpression{}
				} else if parser.mode == CompilerMode {
					panic("The expression for the first argument in a function definition is illegal")
				}
			}
			arguments = append(arguments, firstArgument)
			for parser.match(token.COMMA) {
				a := parser.parseExpression()
				if _, ok := a.(IllegalExpression); ok {
					if parser.mode == ReplMode {
						return IllegalExpression{}
					} else if parser.mode == CompilerMode {
						panic("The expression for an argument in a function definition is illegal")
					}
				}
				arguments = append(arguments, a)
			}
		}
		if t := parser.consumeExpected(token.RPAREN); t.Type == token.ILLEGAL {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("Expected to see a ) in a function defintiion, got something else")
			}
		}
		if t := parser.consumeExpected(token.LBRACE); t.Type == token.ILLEGAL {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("Expected to see a { in a function defintiion, got something else")
			}
		}
	}
	for !parser.match(token.RBRACE) {
		b := parser.parseExpression()
		if _, ok := b.(IllegalExpression); ok {
			if parser.mode == ReplMode {
				return IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("An expression in a function definition block is illegal")
			}
		}
		block = append(block, b)
	}
	return FunctionDefinitionExpression{Type: functionDefinitionExpressionType, Arguments: arguments, Block: block}
}
