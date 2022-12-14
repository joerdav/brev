package tokens

type TokenType string
type Token struct {
	Type     TokenType
	Literal  string
	Col, Row int
}

const (
	IDENT   TokenType = "ident"
	ASSIGN            = "="
	NUMBER            = "number"
	EOF               = "EOF"
	ILLEGAL           = "illegal"

	NE       = "!="
	EQ       = "=="
	ADD      = "+"
	SUB      = "-"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	LT       = "<"
	GT       = ">"
	BANG     = "!"
	LBRK     = "("
	RBRK     = ")"
	LBRC     = "{"
	RBRC     = "}"
	COMMA    = ","

	FUNCTION = "f"
)
