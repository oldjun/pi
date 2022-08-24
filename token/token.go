package token

type Type string

const (
	ILLEGAL           = "ILLEGAL"
	EOF               = "EOF"
	IDENTIFIER        = "IDENTIFIER"
	INT               = "INT"
	FLOAT             = "FLOAT"
	STRING            = "STRING"
	ASSIGN            = "="
	PLUS              = "+"
	PLUS_ASSIGN       = "+="
	PLUS_PLUS         = "++"
	MINUS             = "-"
	MINUS_ASSIGN      = "-="
	MINUS_MINUS       = "--"
	ASTERISK          = "*"
	ASTERISK_ASSIGN   = "*="
	SLASH             = "/"
	SLASH_ASSIGN      = "/="
	MODULO            = "%"
	MODULO_ASSIGN     = "%="
	ASTERISK_ASTERISK = "**"
	BANG              = "!"
	LT                = "<"
	GT                = ">"
	LE                = "<="
	GE                = ">="
	EQ                = "=="
	NE                = "!="
	COMMA             = ","
	SEMICOLON         = ";"
	COLON             = ":"
	LPAREN            = "("
	RPAREN            = ")"
	LBRACE            = "{"
	RBRACE            = "}"
	LBRACKET          = "["
	RBRACKET          = "]"
	DOT               = "."
	TILDE             = "~"
	AT                = "@"
	FUNCTION          = "FUNCTION"
	NULL              = "NULL"
	TRUE              = "TRUE"
	FALSE             = "FALSE"
	IF                = "IF"
	ELSE              = "ELSE"
	RETURN            = "RETURN"
	AND               = "AND"
	OR                = "OR"
	NOT               = "NOT"
	WHILE             = "WHILE"
	BREAK             = "BREAK"
	CONTINUE          = "CONTINUE"
	FOR               = "FOR"
	IN                = "IN"
	QUESTION          = "?"
	BIT_AND           = "&"
	BIT_OR            = "|"
	BIT_XOR           = "^"
	BIT_AND_ASSIGN    = "&="
	BIT_OR_ASSIGN     = "|="
	BIT_XOR_ASSIGN    = "^="
	BIT_RIGHT_SHIFT   = ">>"
	BIT_LEFT_SHIFT    = "<<"
	SWITCH            = "SWITCH"
	CASE              = "CASE"
	DEFAULT           = "DEFAULT"
	CLASS             = "CLASS"
	THIS              = "THIS"
	SUPER             = "SUPER"
	FROM              = "FROM"
	IMPORT            = "IMPORT"
	AS                = "as"
	ASYNC             = "async"
	DEFER             = "defer"
)

var keywords = map[string]Type{
	"func":     FUNCTION,
	"null":     NULL,
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
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"class":    CLASS,
	"this":     THIS,
	"super":    SUPER,
	"from":     FROM,
	"import":   IMPORT,
	"as":       AS,
	"async":    ASYNC,
	"defer":    DEFER,
}

type Token struct {
	Type    Type
	Literal string
	File    string
	Line    int
	Column  int
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
