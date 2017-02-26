package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// ConditionalParslet - the parslet for handling conditionals
type ConditionalParslet struct{}

// ConditionalExpression - an expression representing an if/else statement
type ConditionalExpression struct {
	IfCondition   Expression
	Block         []Expression
	ElseCondition *ConditionalExpression
}

func (e ConditionalExpression) String() string {
	if e.IfCondition != nil {
		if e.ElseCondition == nil {
			return fmt.Sprintf("IF(%s, THEN(%s))", e.IfCondition, expressionListJoin(e.Block, ";"))
		}
		return fmt.Sprintf("IF(%s, THEN(%s), ELSE(%s))", e.IfCondition, expressionListJoin(e.Block, ";"), *e.ElseCondition)
	}
	return expressionListJoin(e.Block, ";")
}

func (p *ConditionalParslet) parse(parser *Parser, inputToken token.Token) Expression {
	block := make([]Expression, 0)
	if parser.match(token.LBRACE) {
		panic(fmt.Sprintf("Missing condition in if statement: %s", inputToken))
	}
	ifCondition := parser.parseExpression()
	parser.match(token.LBRACE)
	for !parser.match(token.RBRACE) {
		block = append(block, parser.parseExpression())
	}
	var elseCondition *ConditionalExpression
	if parser.match(token.ELSE) {
		if parser.match(token.LBRACE) {
			elseBlock := make([]Expression, 0)
			for !parser.match(token.RBRACE) {
				elseBlock = append(elseBlock, parser.parseExpression())
			}
			elseCondition = &ConditionalExpression{Block: elseBlock}
		} else {
			ifToken := parser.consumeExpected(token.IF)
			if elseConditionLiteral, ok := (&ConditionalParslet{}).parse(parser, ifToken).(ConditionalExpression); !ok {
				panic("ConditionalParslet returned a non-conditional expression")
			} else {
				elseCondition = &elseConditionLiteral
			}

		}
	}
	return ConditionalExpression{IfCondition: ifCondition, Block: block, ElseCondition: elseCondition}
}
