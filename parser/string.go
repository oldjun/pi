package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parseString() ast.Expression {
	return &ast.String{Token: p.currToken, Value: p.currToken.Literal}
}
