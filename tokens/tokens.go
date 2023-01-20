package tokens

type (
	TokenType string
	Token     struct {
		Type     TokenType
		Literal  string
		Col, Row int
	}
)

func (tt TokenType) GoString() string {
	switch tt {
	case IDENT:
		return "tokens.IDENT"
	case ASSIGN:
		return "tokens.ASSIGN"
	case NUMBER:
		return "tokens.NUMBER"
	case EOF:
		return "tokens.EOF"
	case ILLEGAL:
		return "tokens.ILLEGAL"
	case NE:
		return "tokens.NE"
	case EQ:
		return "tokens.EQ"
	case ADD:
		return "tokens.ADD"
	case SUB:
		return "tokens.SUB"
	case ASTERISK:
		return "tokens.ASTERISK"
	case SLASH:
		return "tokens.SLASH"
	case PERCENT:
		return "tokens.PERCENT"
	case LT:
		return "tokens.LT"
	case GT:
		return "tokens.GT"
	case BANG:
		return "tokens.BANG"
	case LBRK:
		return "tokens.LBRK"
	case RBRK:
		return "tokens.RBRK"
	case LBRC:
		return "tokens.LBRC"
	case RBRC:
		return "tokens.RBRC"
	case COMMA:
		return "tokens.COMMA"
	case FUNCTION:
		return "tokens.FUNCTION"
	case IF:
		return "tokens.IF"
	case ELSE:
		return "tokens.ELSE"
	case ELIF:
		return "tokens.ELIF"
	case TRUE:
		return "tokens.TRUE"
	case FALSE:
		return "tokens.FALSE"
	default:
		return string(tt)
	}
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

	// Keywords
	FUNCTION = "f"
	IF       = "i"
	ELSE     = "e"
	ELIF     = "ei"
	TRUE     = "T"
	FALSE    = "F"
)
