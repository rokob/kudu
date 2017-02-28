package parser

import (
	"github.com/rokob/kudu/token"
)

// ReplBreakParslet - the parslet for handling the break symbol in the repl
type ReplBreakParslet struct{}

func (p *ReplBreakParslet) parse(parser *Parser, token token.Token) Expression {
	return IllegalExpression{IsBreak: true}
}

func (p *ReplBreakParslet) String() string {
	return "ReplBreakParslet"
}
