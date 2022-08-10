package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
