package parser

import (
	"fmt"
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/lexer"
	"github.com/oldjun/pi/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN  // =,+=,-=,*=,/=,%=,&=,|=,^=
	LOGIC   // &,|,^,&&,||
	TERNARY // ? :
	COMPARE // >,>=,<,<=,==,!=
	SUM     // +
	PRODUCT // *
	PREFIX  // -x
	CALL    // myFunction(X)
	INDEX   // array[index]
	DOT     // some.property or some.function()
)

var precedences = map[token.Type]int{
	token.ASSIGN:          ASSIGN,
	token.PLUS_ASSIGN:     ASSIGN,
	token.MINUS_ASSIGN:    ASSIGN,
	token.ASTERISK_ASSIGN: ASSIGN,
	token.SLASH_ASSIGN:    ASSIGN,
	token.MODULO_ASSIGN:   ASSIGN,
	token.BIT_AND_ASSIGN:  ASSIGN,
	token.BIT_OR_ASSIGN:   ASSIGN,
	token.BIT_XOR_ASSIGN:  ASSIGN,
	token.AND:             LOGIC,
	token.OR:              LOGIC,
	token.IN:              LOGIC,
	token.BIT_AND:         LOGIC,
	token.BIT_OR:          LOGIC,
	token.BIT_XOR:         LOGIC,
	token.BIT_LEFT_SHIFT:  LOGIC,
	token.BIT_RIGHT_SHIFT: LOGIC,
	token.QUESTION:        TERNARY,
	token.EQ:              COMPARE,
	token.NE:              COMPARE,
	token.LT:              COMPARE,
	token.LE:              COMPARE,
	token.GT:              COMPARE,
	token.GE:              COMPARE,
	token.PLUS:            SUM,
	token.MINUS:           SUM,
	token.SLASH:           PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.MODULO:          PRODUCT,
	token.TILDE:           PREFIX,
	token.LPAREN:          CALL,
	token.LBRACKET:        INDEX,
	token.DOT:             DOT,
}

type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l               *lexer.Lexer
	errors          []string
	currToken       token.Token
	peekToken       token.Token
	prefixParseFns  map[token.Type]prefixParseFn
	infixParseFns   map[token.Type]infixParseFn
	postfixParseFns map[token.Type]postfixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParseFns = make(map[token.Type]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.NULL, p.parseNull)
	p.registerPrefix(token.INT, p.parseInteger)
	p.registerPrefix(token.FLOAT, p.parseFloat)
	p.registerPrefix(token.BANG, p.parsePrefix)
	p.registerPrefix(token.TILDE, p.parsePrefix)
	p.registerPrefix(token.NOT, p.parsePrefix)
	p.registerPrefix(token.MINUS, p.parsePrefix)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroup)
	p.registerPrefix(token.LBRACKET, p.parseList)
	p.registerPrefix(token.LBRACE, p.parseHash)
	p.registerPrefix(token.ASTERISK, p.parsePrefix)
	p.registerPrefix(token.ASTERISK_ASTERISK, p.parsePrefix)
	p.registerPrefix(token.IF, p.parseIf)
	p.registerPrefix(token.WHILE, p.parseWhile)
	p.registerPrefix(token.FOR, p.parseFor)
	p.registerPrefix(token.SWITCH, p.parseSwitch)
	p.registerPrefix(token.FUNCTION, p.parseFunction)
	p.registerPrefix(token.STRING, p.parseString)
	p.registerPrefix(token.AT, p.parseDecorator)
	p.registerPrefix(token.CLASS, p.parseClass)
	p.registerPrefix(token.THIS, p.parseThis)
	p.registerPrefix(token.SUPER, p.parseSuper)
	p.registerPrefix(token.FROM, p.parseFrom)
	p.registerPrefix(token.IMPORT, p.parseImport)

	p.infixParseFns = make(map[token.Type]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfix)
	p.registerInfix(token.MINUS, p.parseInfix)
	p.registerInfix(token.SLASH, p.parseInfix)
	p.registerInfix(token.ASTERISK, p.parseInfix)
	p.registerInfix(token.MODULO, p.parseInfix)
	p.registerInfix(token.ASSIGN, p.parseAssign)
	p.registerInfix(token.PLUS_ASSIGN, p.parseCompound)
	p.registerInfix(token.MINUS_ASSIGN, p.parseCompound)
	p.registerInfix(token.SLASH_ASSIGN, p.parseCompound)
	p.registerInfix(token.ASTERISK_ASSIGN, p.parseCompound)
	p.registerInfix(token.MODULO_ASSIGN, p.parseCompound)
	p.registerInfix(token.BIT_AND_ASSIGN, p.parseCompound)
	p.registerInfix(token.BIT_OR_ASSIGN, p.parseCompound)
	p.registerInfix(token.BIT_XOR_ASSIGN, p.parseCompound)
	p.registerInfix(token.EQ, p.parseInfix)
	p.registerInfix(token.NE, p.parseInfix)
	p.registerInfix(token.LT, p.parseInfix)
	p.registerInfix(token.GT, p.parseInfix)
	p.registerInfix(token.LE, p.parseInfix)
	p.registerInfix(token.GE, p.parseInfix)
	p.registerInfix(token.IN, p.parseInfix)
	p.registerInfix(token.LPAREN, p.parseCall)
	p.registerInfix(token.LBRACKET, p.parseIndex)
	p.registerInfix(token.AND, p.parseInfix)
	p.registerInfix(token.OR, p.parseInfix)
	p.registerInfix(token.BIT_AND, p.parseInfix)
	p.registerInfix(token.BIT_OR, p.parseInfix)
	p.registerInfix(token.BIT_XOR, p.parseInfix)
	p.registerInfix(token.BIT_LEFT_SHIFT, p.parseInfix)
	p.registerInfix(token.BIT_RIGHT_SHIFT, p.parseInfix)
	p.registerInfix(token.QUESTION, p.parseTernary)
	p.registerInfix(token.DOT, p.parseDot)

	p.postfixParseFns = make(map[token.Type]postfixParseFn)
	p.registerPostfix(token.PLUS_PLUS, p.parsePostfix)
	p.registerPostfix(token.MINUS_MINUS, p.parsePostfix)

	// Read two tokens, so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currTokenIs(t token.Type) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.Type, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.Type) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}
