package parser

import (
	"fmt"
	"strconv"

	"github.com/rokob/kudu/token"
)

// IntegerParslet - the parslet for handling integers
type IntegerParslet struct{}

// IntegerExpression - an expression representing an integer
type IntegerExpression struct {
	Integer int64
}

func (e IntegerExpression) String() string {
	return fmt.Sprintf("INT(%d)", e.Integer)
}

func (p *IntegerParslet) parse(parser *Parser, token token.Token) Expression {
	i, err := strconv.ParseInt(token.Literal, 0, 64)
	if err != nil {
		if parser.mode == ReplMode {
			return IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic(fmt.Sprintf("Bad Integer value: %s", token))
		}
	}
	return IntegerExpression{Integer: i}
}

func (p *IntegerParslet) String() string {
	return "IntegerParslet"
}
