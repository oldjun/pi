package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseNull() ast.Expression {
	expr := &ast.Null{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return expr
}
