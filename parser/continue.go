package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
