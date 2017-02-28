package token

import "fmt"

// Type is an identifier for each of the different possible Tokens in the language
type Type string

// Token is an object that has lexical meaning in the language
type Token struct {
	Type    Type
	Literal string
	Line    int
	Column  int
}

func (t Token) String() string {
	return fmt.Sprintf("(type: %s, %s)(line: %d, col: %d)", t.Type, t.Literal, t.Line, t.Column)
}

var keywords = map[string]Type{
	"fun":    FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifer - given a string looks up the Type of that identifier
func LookupIdentifer(identifer string) Type {
	if tok, ok := keywords[identifer]; ok {
		return tok
	}
	return IDENT
}

const (
	// ILLEGAL represents a token we don't understand
	ILLEGAL = "ILLEGAL"
	// EOF tells us when to stop lexing
	EOF = "EOF"
	// IDENT - abc, foobar
	IDENT = "IDENT"
	// INT - 34813
	INT = "INT"

	// ASSIGN - the assignment operator
	ASSIGN = "="
	// PLUS - the addition operator
	PLUS = "+"

	// TIMES - the multiplication operator
	TIMES = "*"
	// DIV - the division operator
	DIV = "/"

	// EXP - the exponentiation operator
	EXP = "^"

	// MINUS - the (pre and in)fix minus operator
	MINUS = "-"
	// BANG - the negation prefix operator
	BANG = "!"

	// COMMA - a literal comma
	COMMA = ","
	// SEMICOLON - a literal semicolon
	SEMICOLON = ";"
	// LPAREN - a literal left parenthesis
	LPAREN = "("
	// RPAREN - a literal right parenthesis
	RPAREN = ")"
	// LBRACE - a literal left brace
	LBRACE = "{"
	// RBRACE - a literal left brace
	RBRACE = "}"

	// FUNCTION - keyword for a function
	FUNCTION = "FUNCTION"
	// LET - keyword for declaring a variable
	LET = "LET"
	// IF - keyword for a conditional block
	IF = "IF"
	// ELSE - keyword for an else in a conditional block
	ELSE = "ELSE"
	// RETURN - keyword for returning from a function
	RETURN = "RETURN"

	// DOLLAR - $ symbol used by the interpreter
	DOLLAR = "$"
)
