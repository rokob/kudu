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
		arguments = append(arguments, parser.parseExpression())
		for parser.match(token.COMMA) {
			arguments = append(arguments, parser.parseExpression())
		}
	}
	parser.consumeExpected(token.RPAREN)
	return FunctionCallExpression{Function: left, Arguments: arguments}
}

func (p *FunctionCallParslet) getPrecedence() Precedence {
	return CALL
}
