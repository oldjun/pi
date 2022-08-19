package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parsePostfix(left ast.Expression) ast.Expression {
	expression := &ast.Postfix{
		Token:    p.currToken,
		Left:     left,
		Operator: p.currToken.Literal,
	}
	return expression
}
