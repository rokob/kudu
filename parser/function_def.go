package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// FunctionDefinitionParslet - the parslet for handling function definitions
type FunctionDefinitionParslet struct{}

// FunctionDefinitionExpression - an expression representing a function definition
type FunctionDefinitionExpression struct {
	Arguments []Expression
	Block     []Expression
}

func (e FunctionDefinitionExpression) String() string {
	return fmt.Sprintf("FUN((%s);(%s))", expressionListJoin(e.Arguments, ","), expressionListJoin(e.Block, ";"))
}

func (p *FunctionDefinitionParslet) parse(parser *Parser, inputToken token.Token) Expression {
	arguments := make([]Expression, 0)
	block := make([]Expression, 0)
	if !parser.match(token.LBRACE) {
		parser.consumeExpected(token.LPAREN)
		if !parser.match(token.RPAREN) {
			arguments = append(arguments, parser.parseExpression())
			for parser.match(token.COMMA) {
				arguments = append(arguments, parser.parseExpression())
			}
		}
		parser.consumeExpected(token.RPAREN)
		parser.consumeExpected(token.LBRACE)
	}
	for !parser.match(token.RBRACE) {
		block = append(block, parser.parseExpression())
	}
	return FunctionDefinitionExpression{Arguments: arguments, Block: block}
}
