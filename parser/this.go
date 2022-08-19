package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parseThis() ast.Expression {
	return &ast.This{Token: p.currToken}
}
