package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
