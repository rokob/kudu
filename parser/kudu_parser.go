package parser

import (
	"github.com/rokob/kudu/ast"
	"github.com/rokob/kudu/lexer"
	"github.com/rokob/kudu/token"
)

// ParsingMode is a type which represents the mode this parser is configured to use
type ParsingMode int

const (
	// ReplMode is for the repl which doesn't panic on errors and can handle partial input
	ReplMode ParsingMode = iota
	// CompilerMode expects a fully formed program to be passed to Parse
	CompilerMode
)

// KuduParser is the parser for the kudu language
type KuduParser struct {
	Mode   ParsingMode
	parser *Parser
}

// New - returns an initalized parser for the kudu language
func New(mode ParsingMode) *KuduParser {
	return &KuduParser{Mode: mode}
}

// Parse - parse the input of the kudu language into an Expression.
// Returns (input is legal, input is a repl break, the parsed expressions)
func (p *KuduParser) Parse(input string) (bool, bool, []ast.Expression) {
	lex := lexer.New(input)
	p.parser = NewParser(lex, p.Mode)
	p.configureLanguage()
	parsedExpressions := make([]ast.Expression, 0)
	for {
		parsedExpression := p.parser.parseExpression()
		parsedExpressions = append(parsedExpressions, parsedExpression)

		illegalExp, isIllegal := parsedExpression.(ast.IllegalExpression)
		isBreak := false
		if isIllegal {
			isBreak = illegalExp.IsBreak
			return !isIllegal, isBreak, parsedExpressions
		}
		if !p.parser.match(token.SEMICOLON) {
			return true, false, parsedExpressions
		}
		if p.parser.match(token.EOF) {
			return !isIllegal, isBreak, parsedExpressions
		}
	}
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
		leftCondition: func(e ast.Expression) (bool, string) {
			_, ok := e.(ast.IdentifierExpression)
			return ok, "Left-hand-side of assignment must be an identifier"
		}})

	// (Expression)
	p.parser.registerPrefix(token.LPAREN, &GroupParslet{})

	// a()
	// a(b, c d)
	p.parser.registerInfix(token.LPAREN, &FunctionCallParslet{})

	if p.Mode == ReplMode {
		p.parser.registerPrefix(token.DOLLAR, &ReplBreakParslet{})
	}
}

func prefix(parser *Parser, tokenType token.Type) {
	parser.registerPrefix(tokenType, &PrefixOperatorParslet{})
}
