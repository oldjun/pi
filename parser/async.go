package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseAsync() *ast.Async {
	stmt := &ast.Async{Token: p.currToken}
	p.nextToken()
	stmt.Call = p.parseExpression(LOWEST)
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
