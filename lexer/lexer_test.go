package lexer

import (
	"strings"
	"testing"

	"github.com/joerdav/brev/tokens"
)

func TestNextToken(t *testing.T) {
	input := `
num = 1
num
num = num + 1
num = num - 1
other = num == 1
other = num != 1
1

afunc = f() int {
	a = 1+1
	a
}
bfunc = f(a int) int {
	a+1
}
cfunc = f(a, b int) int {
	a+b
}
dfunc = f(a, b int) {
	a=a+b
}

!1
2*3
2/3
2%3
2<3>4
`
	r := strings.NewReader(input)
	l := NewLexer(r)

	tests := []tokens.Token{
		{Type: tokens.IDENT, Literal: "num", Col: 0, Row: 1},
		{Type: tokens.ASSIGN, Literal: "=", Col: 4, Row: 1},
		{Type: tokens.NUMBER, Literal: "1", Col: 6, Row: 1},

		{Type: tokens.IDENT, Literal: "num", Col: 0, Row: 2},

		{Type: tokens.IDENT, Literal: "num", Col: 0, Row: 3},
		{Type: tokens.ASSIGN, Literal: "=", Col: 4, Row: 3},
		{Type: tokens.IDENT, Literal: "num", Col: 6, Row: 3},
		{Type: tokens.ADD, Literal: "+", Col: 10, Row: 3},
		{Type: tokens.NUMBER, Literal: "1", Col: 12, Row: 3},

		{Type: tokens.IDENT, Literal: "num", Col: 0, Row: 4},
		{Type: tokens.ASSIGN, Literal: "=", Col: 4, Row: 4},
		{Type: tokens.IDENT, Literal: "num", Col: 6, Row: 4},
		{Type: tokens.SUB, Literal: "-", Col: 10, Row: 4},
		{Type: tokens.NUMBER, Literal: "1", Col: 12, Row: 4},

		{Type: tokens.IDENT, Literal: "other", Col: 0, Row: 5},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 5},
		{Type: tokens.IDENT, Literal: "num", Col: 8, Row: 5},
		{Type: tokens.EQ, Literal: "==", Col: 12, Row: 5},
		{Type: tokens.NUMBER, Literal: "1", Col: 15, Row: 5},

		{Type: tokens.IDENT, Literal: "other", Col: 0, Row: 6},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 6},
		{Type: tokens.IDENT, Literal: "num", Col: 8, Row: 6},
		{Type: tokens.NE, Literal: "!=", Col: 12, Row: 6},
		{Type: tokens.NUMBER, Literal: "1", Col: 15, Row: 6},

		{Type: tokens.NUMBER, Literal: "1", Col: 0, Row: 7},

		{Type: tokens.IDENT, Literal: "afunc", Col: 0, Row: 9},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 9},
		{Type: tokens.FUNCTION, Literal: "f", Col: 8, Row: 9},
		{Type: tokens.LBRK, Literal: "(", Col: 9, Row: 9},
		{Type: tokens.RBRK, Literal: ")", Col: 10, Row: 9},
		{Type: tokens.IDENT, Literal: "int", Col: 12, Row: 9},
		{Type: tokens.LBRC, Literal: "{", Col: 16, Row: 9},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 10},
		{Type: tokens.ASSIGN, Literal: "=", Col: 3, Row: 10},
		{Type: tokens.NUMBER, Literal: "1", Col: 5, Row: 10},
		{Type: tokens.ADD, Literal: "+", Col: 6, Row: 10},
		{Type: tokens.NUMBER, Literal: "1", Col: 7, Row: 10},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 11},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 12},

		{Type: tokens.IDENT, Literal: "bfunc", Col: 0, Row: 13},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 13},
		{Type: tokens.FUNCTION, Literal: "f", Col: 8, Row: 13},
		{Type: tokens.LBRK, Literal: "(", Col: 9, Row: 13},
		{Type: tokens.IDENT, Literal: "a", Col: 10, Row: 13},
		{Type: tokens.IDENT, Literal: "int", Col: 12, Row: 13},
		{Type: tokens.RBRK, Literal: ")", Col: 15, Row: 13},
		{Type: tokens.IDENT, Literal: "int", Col: 17, Row: 13},
		{Type: tokens.LBRC, Literal: "{", Col: 21, Row: 13},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 14},
		{Type: tokens.ADD, Literal: "+", Col: 2, Row: 14},
		{Type: tokens.NUMBER, Literal: "1", Col: 3, Row: 14},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 15},

		{Type: tokens.IDENT, Literal: "cfunc", Col: 0, Row: 16},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 16},
		{Type: tokens.FUNCTION, Literal: "f", Col: 8, Row: 16},
		{Type: tokens.LBRK, Literal: "(", Col: 9, Row: 16},
		{Type: tokens.IDENT, Literal: "a", Col: 10, Row: 16},
		{Type: tokens.COMMA, Literal: ",", Col: 11, Row: 16},
		{Type: tokens.IDENT, Literal: "b", Col: 13, Row: 16},
		{Type: tokens.IDENT, Literal: "int", Col: 15, Row: 16},
		{Type: tokens.RBRK, Literal: ")", Col: 18, Row: 16},
		{Type: tokens.IDENT, Literal: "int", Col: 20, Row: 16},
		{Type: tokens.LBRC, Literal: "{", Col: 24, Row: 16},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 17},
		{Type: tokens.ADD, Literal: "+", Col: 2, Row: 17},
		{Type: tokens.IDENT, Literal: "b", Col: 3, Row: 17},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 18},

		{Type: tokens.IDENT, Literal: "dfunc", Col: 0, Row: 19},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 19},
		{Type: tokens.FUNCTION, Literal: "f", Col: 8, Row: 19},
		{Type: tokens.LBRK, Literal: "(", Col: 9, Row: 19},
		{Type: tokens.IDENT, Literal: "a", Col: 10, Row: 19},
		{Type: tokens.COMMA, Literal: ",", Col: 11, Row: 19},
		{Type: tokens.IDENT, Literal: "b", Col: 13, Row: 19},
		{Type: tokens.IDENT, Literal: "int", Col: 15, Row: 19},
		{Type: tokens.RBRK, Literal: ")", Col: 18, Row: 19},
		{Type: tokens.LBRC, Literal: "{", Col: 20, Row: 19},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 20},
		{Type: tokens.ASSIGN, Literal: "=", Col: 2, Row: 20},
		{Type: tokens.IDENT, Literal: "a", Col: 3, Row: 20},
		{Type: tokens.ADD, Literal: "+", Col: 4, Row: 20},
		{Type: tokens.IDENT, Literal: "b", Col: 5, Row: 20},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 21},

		{Type: tokens.BANG, Literal: "!", Col: 0, Row: 23},
		{Type: tokens.NUMBER, Literal: "1", Col: 1, Row: 23},

		{Type: tokens.NUMBER, Literal: "2", Col: 0, Row: 24},
		{Type: tokens.ASTERISK, Literal: "*", Col: 1, Row: 24},
		{Type: tokens.NUMBER, Literal: "3", Col: 2, Row: 24},

		{Type: tokens.NUMBER, Literal: "2", Col: 0, Row: 25},
		{Type: tokens.SLASH, Literal: "/", Col: 1, Row: 25},
		{Type: tokens.NUMBER, Literal: "3", Col: 2, Row: 25},

		{Type: tokens.NUMBER, Literal: "2", Col: 0, Row: 26},
		{Type: tokens.PERCENT, Literal: "%", Col: 1, Row: 26},
		{Type: tokens.NUMBER, Literal: "3", Col: 2, Row: 26},

		{Type: tokens.NUMBER, Literal: "2", Col: 0, Row: 27},
		{Type: tokens.LT, Literal: "<", Col: 1, Row: 27},
		{Type: tokens.NUMBER, Literal: "3", Col: 2, Row: 27},
		{Type: tokens.GT, Literal: ">", Col: 3, Row: 27},
		{Type: tokens.NUMBER, Literal: "4", Col: 4, Row: 27},

		{Type: tokens.EOF, Literal: "", Col: 5, Row: 27},
	}

	for _, tok := range tests {
		c := l.NextToken()
		if c.Type != tok.Type {
			t.Fatalf("c.Type for %s was not the expected value. want=%#v got=%#v", tok.Literal, tok, c)
		}
		if c.Literal != tok.Literal {
			t.Fatalf("c.Literal was not the expected value. want=%#v got=%#v", tok, c)
		}
		if c.Col != tok.Col {
			t.Fatalf("c.Col for %s was not the expected value. want=%#v got=%#v", tok.Literal, tok, c)
		}
		if c.Row != tok.Row {
			t.Fatalf("c.Row for %s was not the expected value. want=%#v got=%#v", tok.Literal, tok, c)
		}
	}
}
