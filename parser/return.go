package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseReturn() *ast.Return {
	stmt := &ast.Return{Token: p.currToken}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
