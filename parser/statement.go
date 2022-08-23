package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.RETURN:
		return p.parseReturn()
	case token.BREAK:
		return p.parseBreak()
	case token.CONTINUE:
		return p.parseContinue()
	case token.ASYNC:
		return p.parseAsync()
	case token.DEFER:
		return p.parseDefer()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
