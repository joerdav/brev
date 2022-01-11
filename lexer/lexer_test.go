package lexer

import (
	"strings"
	"testing"

	"gitlab.com/jrdav/brev/tokens"
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

afunc = fn() int {
	a = 1+1
	return a
}
bfunc = fn(a int) int {
	return a+1
}
cfunc = fn(a, b int) int {
	return a+b
}
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
		{Type: tokens.FUNCTION, Literal: "fn", Col: 8, Row: 9},
		{Type: tokens.LBRK, Literal: "(", Col: 10, Row: 9},
		{Type: tokens.RBRK, Literal: ")", Col: 11, Row: 9},
		{Type: tokens.IDENT, Literal: "int", Col: 13, Row: 9},
		{Type: tokens.LBRC, Literal: "{", Col: 17, Row: 9},
		{Type: tokens.IDENT, Literal: "a", Col: 1, Row: 10},
		{Type: tokens.ASSIGN, Literal: "=", Col: 3, Row: 10},
		{Type: tokens.NUMBER, Literal: "1", Col: 5, Row: 10},
		{Type: tokens.ADD, Literal: "+", Col: 6, Row: 10},
		{Type: tokens.NUMBER, Literal: "1", Col: 7, Row: 10},
		{Type: tokens.IDENT, Literal: "return", Col: 1, Row: 11},
		{Type: tokens.IDENT, Literal: "a", Col: 8, Row: 11},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 12},

		{Type: tokens.IDENT, Literal: "bfunc", Col: 0, Row: 13},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 13},
		{Type: tokens.FUNCTION, Literal: "fn", Col: 8, Row: 13},
		{Type: tokens.LBRK, Literal: "(", Col: 10, Row: 13},
		{Type: tokens.IDENT, Literal: "a", Col: 11, Row: 13},
		{Type: tokens.IDENT, Literal: "int", Col: 13, Row: 13},
		{Type: tokens.RBRK, Literal: ")", Col: 16, Row: 13},
		{Type: tokens.IDENT, Literal: "int", Col: 18, Row: 13},
		{Type: tokens.LBRC, Literal: "{", Col: 22, Row: 13},
		{Type: tokens.IDENT, Literal: "return", Col: 1, Row: 14},
		{Type: tokens.IDENT, Literal: "a", Col: 8, Row: 14},
		{Type: tokens.ADD, Literal: "+", Col: 9, Row: 14},
		{Type: tokens.NUMBER, Literal: "1", Col: 10, Row: 14},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 15},

		{Type: tokens.IDENT, Literal: "cfunc", Col: 0, Row: 16},
		{Type: tokens.ASSIGN, Literal: "=", Col: 6, Row: 16},
		{Type: tokens.FUNCTION, Literal: "fn", Col: 8, Row: 16},
		{Type: tokens.LBRK, Literal: "(", Col: 10, Row: 16},
		{Type: tokens.IDENT, Literal: "a", Col: 11, Row: 16},
		{Type: tokens.COMMA, Literal: ",", Col: 12, Row: 16},
		{Type: tokens.IDENT, Literal: "b", Col: 14, Row: 16},
		{Type: tokens.IDENT, Literal: "int", Col: 16, Row: 16},
		{Type: tokens.RBRK, Literal: ")", Col: 19, Row: 16},
		{Type: tokens.IDENT, Literal: "int", Col: 21, Row: 16},
		{Type: tokens.LBRC, Literal: "{", Col: 25, Row: 16},
		{Type: tokens.IDENT, Literal: "return", Col: 1, Row: 17},
		{Type: tokens.IDENT, Literal: "a", Col: 8, Row: 17},
		{Type: tokens.ADD, Literal: "+", Col: 9, Row: 17},
		{Type: tokens.IDENT, Literal: "b", Col: 10, Row: 17},
		{Type: tokens.RBRC, Literal: "}", Col: 0, Row: 18},

		{Type: tokens.EOF, Literal: "", Col: 1, Row: 18},
	}

	for _, tok := range tests {
		c := l.NextToken()
		if c.Type != tok.Type {
			t.Errorf("c.Type for %s was not the expected value. want=%s got=%s", tok.Literal, tok.Type, c.Type)
		}
		if c.Literal != tok.Literal {
			t.Errorf("c.Literal was not the expected value. want=%s got=%s", tok.Literal, c.Literal)
		}
		if c.Col != tok.Col {
			t.Errorf("c.Col for %s was not the expected value. want=%d got=%d", tok.Literal, tok.Col, c.Col)
		}
		if c.Row != tok.Row {
			t.Errorf("c.Row for %s was not the expected value. want=%d got=%d", tok.Literal, tok.Row, c.Row)
		}
	}
}
