package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseNull() ast.Expression {
	expr := &ast.Null{Token: p.currToken}
	for p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return expr
}
