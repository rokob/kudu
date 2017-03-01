package parser

import (
	"fmt"
	"strconv"

	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// IntegerParslet - the parslet for handling integers
type IntegerParslet struct{}

func (p *IntegerParslet) parse(parser *Parser, token token.Token) ast.Expression {
	i, err := strconv.ParseInt(token.Literal, 0, 64)
	if err != nil {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic(fmt.Sprintf("Bad Integer value: %s", token))
		}
	}
	return ast.IntegerExpression{Integer: i}
}

func (p *IntegerParslet) String() string {
	return "IntegerParslet"
}
