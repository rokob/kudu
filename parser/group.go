package parser

import (
	"github.com/rokob/kudu/token"
)

// GroupParslet - the parselt for handling prefix operators
type GroupParslet struct{}

func (p *GroupParslet) parse(parser *Parser, inputToken token.Token) Expression {
	expression := parser.parseExpression()
	if _, ok := expression.(IllegalExpression); ok {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("The expression inside a group of parentheses is illegal")
		}
	}
	if t := parser.consumeExpected(token.RPAREN); t.Type == token.ILLEGAL {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic("Expected to see a ) in a grouping of expressions, got something else")
		}
	}
	return expression
}

func (p *GroupParslet) String() string {
	return "GroupParslet"
}
