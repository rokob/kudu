package parser

import (
	"fmt"

	"github.com/rokob/kudu/token"
)

// ReturnParslet - the parselt for handling prefix operators
type ReturnParslet struct{}

// ReturnExpression - an expression representing a return from a function
type ReturnExpression struct {
	Value Expression
}

func (e ReturnExpression) String() string {
	return fmt.Sprintf("RETURN(%s)", e.Value)
}

func (p *ReturnParslet) parse(parser *Parser, inputToken token.Token) Expression {
	expression := parser.parseExpression()
	if _, ok := expression.(IllegalExpression); ok {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("An expression being returnd is illegal")
		}
	}
	return ReturnExpression{Value: expression}
}

func (p *ReturnParslet) String() string {
	return "ReturnParslet"
}
