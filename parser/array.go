package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseArray() ast.Expression {
	array := &ast.Array{Token: p.currToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}
