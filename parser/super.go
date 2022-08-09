package parser

import "pilang/ast"

func (p *Parser) parseSuper() ast.Expression {
	return &ast.Super{Token: p.currToken}
}
