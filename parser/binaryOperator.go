package parser

import (
	"fmt"

	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/token"
)

// BinaryOperatorParselet - parslet for binary operators
type BinaryOperatorParselet struct {
	precedence     Precedence
	isRight        bool
	leftCondition  func(ast.Expression) (bool, string)
	rightCondition func(ast.Expression) (bool, string)
}

func (p *BinaryOperatorParselet) parse(parser *Parser, left ast.Expression, token token.Token) ast.Expression {
	precedence := p.precedence
	if p.isRight {
		precedence--
	}
	if p.leftCondition != nil {
		if ok, msg := p.leftCondition(left); !ok {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic(msg)
			}
		}
	}
	right := parser.parseExpression(precedence)
	if _, ok := right.(ast.IllegalExpression); ok {
		if parser.mode == ReplMode {
			return ast.IllegalExpression{}
		} else if parser.mode == CompilerMode {
			panic(fmt.Sprintf("The condition on the right side of the binary operator %s is illegal", token.Literal))
		}
	}
	if p.rightCondition != nil {
		if ok, msg := p.rightCondition(right); !ok {
			if parser.mode == ReplMode {
				return ast.IllegalExpression{}
			} else if parser.mode == CompilerMode {
				panic(msg)
			}
		}
	}
	return ast.OperatorExpression{
		Type: token.Type,
		Arguments: ast.OperatorExpressionArguments{
			Left:  left,
			Right: right,
		},
	}
}

func (p *BinaryOperatorParselet) getPrecedence() Precedence {
	return p.precedence
}

func (p *BinaryOperatorParselet) String() string {
	return fmt.Sprintf("BinaryOperatorParselet(precedence: %d, isRight: %t)", p.precedence, p.isRight)
}
