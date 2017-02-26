package lexer

import "github.com/rokob/kudu/token"

// Lexer is an object that process an input string and emits a stream of tokens
type Lexer struct {
	input         string
	position      int // points to place in input where char is
	readPosition  int // points to place in input where we will read next
	char          byte
	currentLine   int
	currentColumn int
}

// New - returns a reference to a newly constructed lexer for the given input string
func New(input string) *Lexer {
	l := &Lexer{input: input, currentLine: 1}
	l.readChar()
	return l
}

// NextToken - consumes enough of the input to find the next token which it emits
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tok = l.newCurrentToken(token.ASSIGN)
	case ';':
		tok = l.newCurrentToken(token.SEMICOLON)
	case '+':
		tok = l.newCurrentToken(token.PLUS)
	case '-':
		tok = l.newCurrentToken(token.MINUS)
	case '*':
		tok = l.newCurrentToken(token.TIMES)
	case '/':
		tok = l.newCurrentToken(token.DIV)
	case '^':
		tok = l.newCurrentToken(token.EXP)
	case '(':
		tok = l.newCurrentToken(token.LPAREN)
	case ')':
		tok = l.newCurrentToken(token.RPAREN)
	case '{':
		tok = l.newCurrentToken(token.LBRACE)
	case '}':
		tok = l.newCurrentToken(token.RBRACE)
	case ',':
		tok = l.newCurrentToken(token.COMMA)
	case '!':
		tok = l.newCurrentToken(token.BANG)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal, tok.Line, tok.Column = l.readIdentifier()
			tok.Type = token.LookupIdentifer(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal, tok.Line, tok.Column = l.readInteger()
			return tok
		}
		tok = l.newCurrentToken(token.ILLEGAL)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.updateSourcePosition()
}

func (l *Lexer) updateSourcePosition() {
	if isNewline(l.char) {
		l.currentLine++
		l.currentColumn = 0
	} else {
		l.currentColumn++
	}
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() (ident string, line int, column int) {
	return l.readValue(func() {
		for isIdentifierMiddleLegal(l.char) {
			l.readChar()
		}
		if isIdentifierTerminalLegal(l.char) {
			l.readChar()
		}
	})
}

func (l *Lexer) readInteger() (digit string, line int, column int) {
	return l.readValue(func() {
		for isDigit(l.char) {
			l.readChar()
		}
	})
}

func (l *Lexer) readValue(reader func()) (digit string, line int, column int) {
	position := l.position
	line = l.currentLine
	column = l.currentColumn
	reader()
	return l.input[position:l.position], line, column
}

func (l *Lexer) newCurrentToken(tokenType token.Type) token.Token {
	return newToken(tokenType, l.char, l.currentLine, l.currentColumn)
}

func newToken(tokenType token.Type, ch byte, line int, column int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line, Column: column}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isNewline(ch byte) bool {
	return ch == '\n'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isIdentifierMiddleLegal(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}

func isIdentifierTerminalLegal(ch byte) bool {
	return ch == '!' || ch == '?'
}
