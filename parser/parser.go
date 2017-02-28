package parser

import (
	"fmt"
	"strings"

	"github.com/rokob/kudu/token"
)

// TokenStream - interface representing a stream of tokens
type TokenStream interface {
	NextToken() token.Token
}

// Expression - generic expression interface
type Expression interface {
	String() string
}

// PrefixParslet - parses a prefix expression
type PrefixParslet interface {
	parse(parser *Parser, token token.Token) Expression
}

// InfixParslet - parses an infix expression
type InfixParslet interface {
	parse(parser *Parser, left Expression, token token.Token) Expression
	getPrecedence() Precedence
}

// Parser - An expression parser
type Parser struct {
	read           []token.Token
	stream         TokenStream
	prefixParslets map[token.Type]PrefixParslet
	infixParslets  map[token.Type]InfixParslet
	mode           ParsingMode
}

// NewParser - returns a properly initialized parser reference
func NewParser(tokenStream TokenStream, mode ParsingMode) *Parser {
	p := Parser{
		read:           make([]token.Token, 0),
		stream:         tokenStream,
		prefixParslets: make(map[token.Type]PrefixParslet),
		infixParslets:  make(map[token.Type]InfixParslet),
		mode:           mode,
	}
	return &p
}

func (p *Parser) parseExpression(precedenceParam ...Precedence) Expression {
	var precedence Precedence
	if len(precedenceParam) == 0 {
		precedence = LOWEST
	} else if len(precedenceParam) == 1 {
		precedence = precedenceParam[0]
	} else {
		if p.mode == ReplMode {
			return IllegalExpression{}
		} else if p.mode == CompilerMode {
			panic(fmt.Sprintf("Too many parameters to parseExpression: %v", precedenceParam))
		}
	}

	token := p.consume()
	prefix, ok := p.prefixParslets[token.Type]
	if !ok {
		if p.mode == ReplMode {
			return IllegalExpression{}
		} else if p.mode == CompilerMode {
			panic(fmt.Sprintf("Bad token: %s", token))
		}
	}
	left := prefix.parse(p, token)

	for precedence < p.getPrecedence() {
		token = p.consume()

		infix, ok := p.infixParslets[token.Type]
		if !ok {
			return left
		}
		left = infix.parse(p, left, token)
	}
	return left
}

func (p *Parser) registerPrefix(tokenType token.Type, parslet PrefixParslet) {
	p.prefixParslets[tokenType] = parslet
}

func (p *Parser) registerInfix(tokenType token.Type, parslet InfixParslet) {
	p.infixParslets[tokenType] = parslet
}

func (p *Parser) lookAhead(distance int) token.Token {
	for distance >= len(p.read) {
		p.read = append(p.read, p.stream.NextToken())
	}
	return p.read[distance]
}

func (p *Parser) consume() token.Token {
	p.lookAhead(0)
	token := p.read[0]
	p.read = p.read[1:]
	return token
}

func (p *Parser) consumeExpected(expected token.Type) token.Token {
	nextToken := p.lookAhead(0)
	if nextToken.Type != expected {
		if p.mode == ReplMode {
			return token.Token{Type: token.ILLEGAL}
		} else if p.mode == CompilerMode {
			panic(fmt.Sprintf("Expected: %s, got: %s", expected, nextToken))
		}
	}
	return p.consume()
}

func (p *Parser) match(expected token.Type) bool {
	token := p.lookAhead(0)
	if token.Type != expected {
		return false
	}
	p.consume()
	return true
}

func (p *Parser) getPrecedence() Precedence {
	nextTokenType := p.lookAhead(0).Type
	if parslet, ok := p.infixParslets[nextTokenType]; ok {
		return parslet.getPrecedence()
	}
	return LOWEST
}

func expressionListJoin(list []Expression, sep string) string {
	expressions := make([]string, len(list))
	for i, e := range list {
		expressions[i] = e.String()
	}
	return strings.Join(expressions, sep)
}
