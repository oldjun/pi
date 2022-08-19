package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parseSuper() ast.Expression {
	return &ast.Super{Token: p.currToken}
}
