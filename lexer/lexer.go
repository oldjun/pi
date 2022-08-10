package lexer

import (
	"pilang/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	file         string
	line         int
	column       int
}

func New(input string, file string) *Lexer {
	l := &Lexer{input: input, file: file, line: 1, column: 1}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ASSIGN, string(l.ch))
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.NE, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.BANG, string(l.ch))
		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.PLUS_ASSIGN, string(ch)+string(l.ch))
		} else if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.PLUS_PLUS, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.PLUS, string(l.ch))
		}
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.MINUS_ASSIGN, string(ch)+string(l.ch))
		} else if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.MINUS_MINUS, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.MINUS, string(l.ch))
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.ASTERISK_ASSIGN, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ASTERISK, string(l.ch))
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.SLASH_ASSIGN, string(ch)+string(l.ch))
		} else if l.peekChar() == '/' {
			_ = l.readLine()
			l.readChar()
			return l.NextToken()
		} else {
			tok = l.newToken(token.SLASH, string(l.ch))
		}
	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.MODULO_ASSIGN, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.MODULO, string(l.ch))
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.LE, string(ch)+string(l.ch))
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.BIT_LEFT_SHIFT, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.LT, string(l.ch))
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.GE, string(ch)+string(l.ch))
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.BIT_RIGHT_SHIFT, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.GT, string(l.ch))
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.AND, string(ch)+string(l.ch))
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.BIT_AND_ASSIGN, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.BIT_AND, string(l.ch))
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.OR, string(ch)+string(l.ch))
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.BIT_OR_ASSIGN, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.BIT_OR, string(l.ch))
		}
	case '^':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.BIT_XOR_ASSIGN, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.BIT_XOR, string(l.ch))
		}
	case ';':
		tok = l.newToken(token.SEMICOLON, string(l.ch))
	case ':':
		tok = l.newToken(token.COLON, string(l.ch))
	case '(':
		tok = l.newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = l.newToken(token.RPAREN, string(l.ch))
	case '~':
		tok = l.newToken(token.TILDE, string(l.ch))
	case ',':
		tok = l.newToken(token.COMMA, string(l.ch))
	case '{':
		tok = l.newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = l.newToken(token.RBRACE, string(l.ch))
	case '[':
		tok = l.newToken(token.LBRACKET, string(l.ch))
	case ']':
		tok = l.newToken(token.RBRACKET, string(l.ch))
	case '@':
		tok = l.newToken(token.AT, string(l.ch))
	case '"':
		tok = l.newToken(token.STRING, l.readString())
	case '?':
		tok = l.newToken(token.QUESTION, string(l.ch))
	case '.':
		tok = l.newToken(token.DOT, string(l.ch))
	case 0:
		tok = l.newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := token.LookupIdent(literal)
			tok = l.newToken(tokenType, literal)
			return tok
		} else if isDigit(l.ch) {
			tok = l.readDecimal()
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType token.Type, literal string) token.Token {
	tok := token.Token{Type: tokenType, Literal: literal, File: l.file, Line: l.line, Column: l.column}
	l.column += len(literal)
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	l.readChar()
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[position:l.position]
}

// Go ahead until you find a new line.
// This makes it so that comments take a full line.
func (l *Lexer) readLine() string {
	position := l.position
	for {
		l.readChar()
		if l.ch == '\n' || l.ch == '\r' || l.ch == 0 {
			break
		}
	}
	l.line += 1
	l.column = 1
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line += 1
			l.column = 1
		} else if l.ch == ' ' {
			l.column += 1
		} else if l.ch == '\t' {
			l.column += 4
		}
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readDecimal() token.Token {
	integer := l.readNumber()
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar()
		fraction := l.readNumber()
		return l.newToken(token.FLOAT, integer+"."+fraction)
	}
	return l.newToken(token.INT, integer)
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
