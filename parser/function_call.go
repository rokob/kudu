package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// FunctionCallParslet - the parslet for handling function calls
type FunctionCallParslet struct{}

// FunctionCallExpression - an expression representing a function call
type FunctionCallExpression struct {
	Function  Expression
	Arguments []Expression
}

func (e FunctionCallExpression) String() string {
	return fmt.Sprintf("CALL(%s;(%s))", e.Function, expressionListJoin(e.Arguments, ","))
}

func (p *FunctionCallParslet) parse(parser *Parser, left Expression, inputToken token.Token) Expression {
	arguments := make([]Expression, 0)
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
			panic("Expected to see a ) in a function call, got something else")
		}
	}

	return FunctionCallExpression{Function: left, Arguments: arguments}
}

func (p *FunctionCallParslet) getPrecedence() Precedence {
	return CALL
}
