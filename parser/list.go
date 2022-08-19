package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseList() ast.Expression {
	array := &ast.List{Token: p.currToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}
