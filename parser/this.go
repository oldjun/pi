package parser

import "pilang/ast"

func (p *Parser) parseThis() ast.Expression {
	return &ast.This{Token: p.currToken}
}
