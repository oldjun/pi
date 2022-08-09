package token

type Type string

const (
	ILLEGAL         = "ILLEGAL"
	EOF             = "EOF"
	IDENTIFIER      = "IDENTIFIER" // add, foobar, x, y, ...
	INT             = "INT"        // 123
	FLOAT           = "FLOAT"      // 1.23
	STRING          = "STRING"
	ASSIGN          = "="
	PLUS            = "+"
	PLUS_ASSIGN     = "+="
	PLUS_PLUS       = "++"
	MINUS           = "-"
	MINUS_ASSIGN    = "-="
	MINUS_MINUS     = "--"
	ASTERISK        = "*"
	ASTERISK_ASSIGN = "*="
	SLASH           = "/"
	SLASH_ASSIGN    = "/="
	MODULO          = "%"
	MODULO_ASSIGN   = "%="
	BANG            = "!"
	LT              = "<"
	GT              = ">"
	LE              = "<="
	GE              = ">="
	EQ              = "=="
	NE              = "!="
	COMMA           = ","
	SEMICOLON       = ";"
	COLON           = ":"
	LPAREN          = "("
	RPAREN          = ")"
	LBRACE          = "{"
	RBRACE          = "}"
	LBRACKET        = "["
	RBRACKET        = "]"
	DOT             = "."
	TILDE           = "~"
	AT              = "@"
	FUNCTION        = "FUNCTION"
	TRUE            = "TRUE"
	FALSE           = "FALSE"
	IF              = "IF"
	ELSE            = "ELSE"
	RETURN          = "RETURN"
	AND             = "AND"
	OR              = "OR"
	NOT             = "NOT"
	WHILE           = "WHILE"
	BREAK           = "BREAK"
	CONTINUE        = "CONTINUE"
	FOR             = "FOR"
	IN              = "IN"
	QUESTION        = "?"
	BIT_AND         = "&"
	BIT_OR          = "|"
	BIT_XOR         = "^"
	BIT_AND_ASSIGN  = "&="
	BIT_OR_ASSIGN   = "|="
	BIT_XOR_ASSIGN  = "^="
	BIT_RIGHT_SHIFT = ">>"
	BIT_LEFT_SHIFT  = "<<"
	CLASS           = "CLASS"
	THIS            = "THIS"
	SUPER           = "SUPER"
)

var keywords = map[string]Type{
	"func":     FUNCTION,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"while":    WHILE,
	"break":    BREAK,
	"continue": CONTINUE,
	"for":      FOR,
	"in":       IN,
	"class":    CLASS,
	"this":     THIS,
	"super":    SUPER,
}

type Token struct {
	Type    Type
	Literal string
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
