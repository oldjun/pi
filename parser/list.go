package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseList() ast.Expression {
	array := &ast.List{Token: p.currToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}
