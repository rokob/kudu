package parser

import (
	"github.com/rokob/kudu/lexer"
	"github.com/rokob/kudu/token"
)

// KuduParser is the parser for the kudu language
type KuduParser struct {
	parser *Parser
}

// New - returns an initalized parser for the kudu language
func New() *KuduParser {
	return &KuduParser{}
}

// Parse - parse the input of the kudu language into an Expression
func (p *KuduParser) Parse(input string) Expression {
	lex := lexer.New(input)
	p.parser = NewParser(lex)
	p.configureLanguage()
	return p.parser.parseExpression()
}

func (p *KuduParser) configureLanguage() {
	// anythingThatIsNotAKeyword
	// canHave_And1234AndCanEndInOne!
	// canHave_And1234OrInOne?
	// FAIL: something-else
	// FAIL: ajd!!
	// FAIL: whatabout!847
	// FAIL: 12ajbhf
	p.parser.registerPrefix(token.IDENT, &IdentifierParslet{})

	// 1234
	// 39049284
	p.parser.registerPrefix(token.INT, &IntegerParslet{})

	// let a = Expresssion
	p.parser.registerPrefix(token.LET, &LetParslet{})

	// if !x { do(x) }
	// if !x { do(x) } else { do(y) }
	// if !x { do(x) } else if z { do(z) }
	// if !x { do(x) } else if z { do(z) } else { do(y) }
	p.parser.registerPrefix(token.IF, &ConditionalParslet{})

	// fun (x,y) { x(y) }
	// fun { doSomething() }
	p.parser.registerPrefix(token.FUNCTION, &FunctionDefinitionParslet{})

	// return Expression
	p.parser.registerPrefix(token.RETURN, &ReturnParslet{})

	// -42
	// -stuff
	prefix(p.parser, token.MINUS)

	// !xyz
	prefix(p.parser, token.BANG)

	// 3 + a
	p.parser.registerInfix(token.PLUS, &BinaryOperatorParselet{precedence: SUM})
	// 5 - 2
	p.parser.registerInfix(token.MINUS, &BinaryOperatorParselet{precedence: SUM})
	// a * b
	p.parser.registerInfix(token.TIMES, &BinaryOperatorParselet{precedence: PRODUCT})
	// 32 / 99
	p.parser.registerInfix(token.DIV, &BinaryOperatorParselet{precedence: PRODUCT})
	// 32 ^ n
	// 2 ^ 3 ^ 4 -> 2 ^ (3 ^ 4)
	p.parser.registerInfix(token.EXP, &BinaryOperatorParselet{precedence: EXPONENT, isRight: true})

	// foo = 32
	// a = b = c -> a = (b = c)
	// FAIL: 32 = a
	p.parser.registerInfix(token.ASSIGN, &BinaryOperatorParselet{
		precedence: ASSIGNMENT,
		isRight:    true,
		leftCondition: func(e Expression) (bool, string) {
			_, ok := e.(IdentifierExpression)
			return ok, "Left-hand-side of assignment must be an identifier"
		}})

	// (Expression)
	p.parser.registerPrefix(token.LPAREN, &GroupParslet{})

	// a()
	// a(b, c d)
	p.parser.registerInfix(token.LPAREN, &FunctionCallParslet{})
}

func prefix(parser *Parser, tokenType token.Type) {
	parser.registerPrefix(tokenType, &PrefixOperatorParslet{})
}
