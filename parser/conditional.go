package parser

import (
	"fmt"

	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// ConditionalParslet - the parslet for handling conditionals
type ConditionalParslet struct{}

func (p *ConditionalParslet) parse(parser *Parser, inputToken token.Token) ast.Expression {
	block := make([]ast.Expression, 0)
	if parser.match(token.LBRACE) {
		panic(fmt.Sprintf("Missing condition in if statement: %s", inputToken))
	}
	ifCondition := parser.parseExpression()
	if _, ok := ifCondition.(ast.IllegalExpression); ok {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("The condition for an if statement is illegal")
		}
	}
	lbraceToken := parser.consumeExpected(token.LBRACE)
	if lbraceToken.Type == token.ILLEGAL {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("After an if statement we expeceted a { but got something else")
		}

	}
	for !parser.match(token.RBRACE) {
		b := parser.parseExpression()
		if _, ok := b.(ast.IllegalExpression); ok {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic("An expression inside an if block is illegal")
			}
		}
		block = append(block, b)
	}
	var elseCondition *ast.ConditionalExpression
	if parser.match(token.ELSE) {
		if parser.match(token.LBRACE) {
			elseBlock := make([]ast.Expression, 0)
			for !parser.match(token.RBRACE) {
				e := parser.parseExpression()
				if _, ok := e.(ast.IllegalExpression); ok {
					if parser.mode == ReplMode {
						return ast.IllegalExpression{}
					} else if parser.mode == CompilerMode {
						panic("An expression inside an else block is illegal")
					}
				}
				elseBlock = append(elseBlock, e)
			}
			elseCondition = &ast.ConditionalExpression{Block: elseBlock}
		} else {
			ifToken := parser.consumeExpected(token.IF)
			if ifToken.Type == token.ILLEGAL {
				if parser.mode == ReplMode {
					return ast.IllegalExpression{}
				} else if parser.mode == CompilerMode {
					panic("Expected to consume an IF token, but instead got something else")
				}
			}
			if elseConditionLiteral, ok := (&ConditionalParslet{}).parse(parser, ifToken).(ast.ConditionalExpression); !ok {
				if parser.mode == ReplMode {
					return ast.IllegalExpression{}
				} else if parser.mode == CompilerMode {
					panic("ConditionalParslet returned a non-conditional expression")
				}

			} else {
				elseCondition = &elseConditionLiteral
			}

		}
	}
	return ast.ConditionalExpression{IfCondition: ifCondition, Block: block, ElseCondition: elseCondition}
}
