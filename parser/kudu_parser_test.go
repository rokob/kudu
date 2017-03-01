package parser

import (
	"testing"
)

func TestKuduParser_Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Simple ident",
			"baz32_bar", "IDENT(baz32_bar)"},
		{"Simple prefix",
			"!foo", "PREFIX(!, IDENT(foo))"},
		{"Simple binary",
			"foo + bar", "BINARY(+, IDENT(foo), IDENT(bar))"},
		{"Prefixes",
			"-!!foo", "PREFIX(-, PREFIX(!, PREFIX(!, IDENT(foo))))"},
		{"Multiple things",
			"!foo - bar + 32",
			"PREFIX(!, BINARY(+, BINARY(-, IDENT(foo), IDENT(bar)), INT(32)))"},
		{"Exponentials",
			"a ^ b ^ c",
			"BINARY(^, IDENT(a), BINARY(^, IDENT(b), IDENT(c)))"},
		{"Function call",
			"a(b,c,d)",
			"CALL(IDENT(a);(IDENT(b),IDENT(c),IDENT(d)))"},
		{"Arithmetic precedence",
			"a + b * 3",
			"BINARY(+, IDENT(a), BINARY(*, IDENT(b), INT(3)))"},
		{"Grouping",
			"(a + b) * 3",
			"BINARY(*, BINARY(+, IDENT(a), IDENT(b)), INT(3))"},
		{"Let binding",
			"let a = 4*b",
			"LET(IDENT(a) = BINARY(*, INT(4), IDENT(b)))"},
		{"Assignment",
			"buzz = fizz",
			"BINARY(=, IDENT(buzz), IDENT(fizz))"},
		{"Multiple assignment",
			"foo = bar = baz",
			"BINARY(=, IDENT(foo), BINARY(=, IDENT(bar), IDENT(baz)))"},
		{"Conditional",
			`if !a {
				b(x)
			}`,
			"IF(PREFIX(!, IDENT(a)), THEN(CALL(IDENT(b);(IDENT(x)))))"},
		{"Conditional with multiple statements",
			`if !a {
				b(x)
				4 + 2
			}`,
			"IF(PREFIX(!, IDENT(a)), THEN(CALL(IDENT(b);(IDENT(x)));BINARY(+, INT(4), INT(2))))"},
		{"Conditional with else",
			`if a {
    			b(x)
			} else {
    			d(x)
			}`,
			"IF(IDENT(a), THEN(CALL(IDENT(b);(IDENT(x)))), ELSE(CALL(IDENT(d);(IDENT(x)))))"},

		{"Conditional with elseif and else",
			`if a {
    			b(x)
			} else if !c {
   				d(x)
			} else {
    			f(x)
			}`,
			"IF(IDENT(a), THEN(CALL(IDENT(b);(IDENT(x)))), ELSE(IF(PREFIX(!, IDENT(c)), THEN(CALL(IDENT(d);(IDENT(x)))), ELSE(CALL(IDENT(f);(IDENT(x)))))))"},
		{"Conditional with elseif and no else",
			`if a {
    			b(x)
			} else if !c {
   				d(x)
			}`,
			"IF(IDENT(a), THEN(CALL(IDENT(b);(IDENT(x)))), ELSE(IF(PREFIX(!, IDENT(c)), THEN(CALL(IDENT(d);(IDENT(x)))))))"},
		{"Function no arguments",
			`fun {
    			return 42
			 }`,
			"FUN(();(RETURN(INT(42))))"},
		{"Function with arguments",
			`fun(x, y) {
    			x(y)
		 	 }`,
			"FUN((IDENT(x),IDENT(y));(CALL(IDENT(x);(IDENT(y)))))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(CompilerMode)
			_, _, gotExpr := p.Parse(tt.input)
			got := gotExpr[0].String()
			if got != tt.want {
				t.Errorf("Parser.parseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKuduParser_ParsePanics(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Non identifier assignment", "42 + 3 = b"},
		{"Malformed input", ",a=2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(CompilerMode)
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected KuduParser.parse with input %v to panic", tt.input)
				}
			}()
			p.Parse(tt.input)
		})
	}
}
