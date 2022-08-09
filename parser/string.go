package parser

import "pilang/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Token: p.currToken, Value: p.currToken.Literal}
}
