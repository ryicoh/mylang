package mylang

import (
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input  string
		expect []token
	}{
		{
			input: "1+2", expect: []token{{INT, "1"}, {ADD, "+"}, {INT, "2"}},
		},
		{
			input: "1-2", expect: []token{{INT, "1"}, {SUB, "-"}, {INT, "2"}},
		},
		{
			input: "1*2", expect: []token{{INT, "1"}, {MUL, "*"}, {INT, "2"}},
		},
		{
			input: "1/2", expect: []token{{INT, "1"}, {DIV, "/"}, {INT, "2"}},
		},
		{
			input: "1.1+2.2", expect: []token{{FLOAT, "1.1"}, {ADD, "+"}, {FLOAT, "2.2"}},
		},
		{
			input: "74-(12+34)", expect: []token{
				{INT, "74"}, {SUB, "-"}, {LPAREN, "("},
				{INT, "12"}, {ADD, "+"}, {INT, "34"}, {RPAREN, ")"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := NewLexer(strings.NewReader(tt.input))

			for _, expect := range tt.expect {
				sym := new(yySymType)
				actual := l.Lex(sym)

				if expect.typ != actual {
					t.Errorf("expect (%v), but got (%v)", expect.typ, actual)
				}

				if !reflect.DeepEqual(expect, sym.token) {
					t.Errorf("expect (%v), but got (%v)", expect, sym.token)
				}
			}
		})
	}
}
