package lexer

import (
	"testing"

	"github.com/rokob/kudu/token"
)

type nextTokenTest struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []nextTokenTest{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	nextTokenHelper(input, tests, t)
}

func TestNextTokenConditional(t *testing.T) {
	input := `if wookie {
        a(b)
    }`
	tests := []nextTokenTest{
		{token.IF, "if"},
		{token.IDENT, "wookie"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.LPAREN, "("},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
	}

	nextTokenHelper(input, tests, t)
}

func TestNextTokenConditionalWithElse(t *testing.T) {
	input := `if wookie {
        a(b)
    } else if b {
        c
    }`
	tests := []nextTokenTest{
		{token.IF, "if"},
		{token.IDENT, "wookie"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.LPAREN, "("},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.IDENT, "b"},
		{token.LBRACE, "{"},
		{token.IDENT, "c"},
		{token.RBRACE, "}"},
	}

	nextTokenHelper(input, tests, t)
}

func TestNextTokenFunctionDef(t *testing.T) {
	input := `fun (x) {
        return a
    }`
	tests := []nextTokenTest{
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.RBRACE, "}"},
	}

	nextTokenHelper(input, tests, t)
}

func TestNextTokenComplex(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fun(x, y) {
x + !y - ten;
};
let result = add(five, ten);
let what? = 37
a(b,c)
`
	tests := []nextTokenTest{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.BANG, "!"},
		{token.IDENT, "y"},
		{token.MINUS, "-"},
		{token.IDENT, "ten"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "what?"},
		{token.ASSIGN, "="},
		{token.INT, "37"},
		{token.IDENT, "a"},
		{token.LPAREN, "("},
		{token.IDENT, "b"},
		{token.COMMA, ","},
		{token.IDENT, "c"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}
	nextTokenHelper(input, tests, t)
}

func nextTokenHelper(input string, tests []nextTokenTest, t *testing.T) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
