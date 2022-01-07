package lexer

import (
	"bufio"
	"io"

	"github.com/joe-davidson1802/brev/tokens"
)

type Lexer struct {
	reader        *bufio.Reader
	current, peek *rune
	row, col      int
}

func NewLexer(reader io.Reader) Lexer {
	t := Lexer{}
	t.reader = bufio.NewReader(reader)
	t.Advance()
	t.Advance()
	t.row = 0
	t.col = 0
	return t
}

var singleRuneTokens = map[rune]tokens.TokenType{
	'=': tokens.ASSIGN,
	'+': tokens.ADD,
	'-': tokens.SUB,
	'{': tokens.LBRC,
	'}': tokens.RBRC,
	'(': tokens.LBRK,
	')': tokens.RBRK,
	',': tokens.COMMA,
}

var keywords = map[string]tokens.TokenType{
	"fn": tokens.FUNCTION,
}

func (t *Lexer) NextToken() tokens.Token {
	if t.current == nil || t.peek == nil {
		return t.token(tokens.EOF, "")
	}
	t.eatWhitespace()
	ty, single := singleRuneTokens[*t.current]
	switch {
	case validIdentFirstChar(*t.current):
		return t.parseIdent()
	case t.isCurrent('!') && t.isPeek('='):
		to := t.token(tokens.NE, "!=")
		t.Advance()
		t.Advance()
		return to
	case t.isCurrent('=') && t.isPeek('='):
		to := t.token(tokens.EQ, "==")
		t.Advance()
		t.Advance()
		return to
	case single:
		to := t.currentAsToken(ty)
		t.Advance()
		return to
	case validNumberChar(*t.current):
		return t.parseNumber()
	default:
		to := t.currentAsToken(tokens.ILLEGAL)
		t.Advance()
		return to
	}
}

func (t *Lexer) token(tt tokens.TokenType, literal string) tokens.Token {
	return tokens.Token{Type: tt, Literal: literal, Col: t.col, Row: t.row}
}

func (t *Lexer) currentAsToken(tt tokens.TokenType) tokens.Token {
	return tokens.Token{Type: tt, Literal: string(*t.current), Col: t.col, Row: t.row}
}

func (t *Lexer) Advance() {
	t.col++
	if t.current != nil && t.isCurrent('\n') || t.isCurrent('\r') {
		t.row++
		t.col = 0
	}
	t.current = t.peek
	r, _, err := t.reader.ReadRune()
	if err != nil {
		t.peek = nil
		return
	}
	t.peek = &r
}

func (t *Lexer) parseIdent() tokens.Token {
	to := t.token(tokens.IDENT, "")
	for t.current != nil && validIdentChar(*t.current) {
		to.Literal += string(*t.current)
		t.Advance()
	}
	if k, ok := keywords[to.Literal]; ok {
		to.Type = k
	}
	return to
}

func (t *Lexer) parseNumber() tokens.Token {
	to := t.token(tokens.NUMBER, "")
	for t.current != nil && validNumberChar(*t.current) {
		to.Literal += string(*t.current)
		t.Advance()
	}
	return to
}

func (t *Lexer) isCurrent(r rune) bool {
	return t.current != nil && *t.current == r
}

func (t *Lexer) isPeek(r rune) bool {
	return t.peek != nil && *t.peek == r
}

func (t *Lexer) eatWhitespace() {
	for t.isCurrent('\n') || t.isCurrent('\r') || t.isCurrent('\t') || t.isCurrent(' ') {
		t.Advance()
	}
}

func validNumberChar(r rune) bool {
	i := uint64(r)
	return uint64('0') <= i && i <= uint64('9')
}

func validIdentFirstChar(r rune) bool {
	i := uint64(r)
	return uint64('a') <= i && i <= uint64('z') || uint64('A') <= i && i <= uint64('Z') || r == rune('_')
}

func validIdentChar(r rune) bool {
	i := uint64(r)
	return validIdentFirstChar(r) || uint64('0') <= i && i <= uint64('9')
}
