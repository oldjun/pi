package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.currToken, Value: p.curTokenIs(token.TRUE)}
}
