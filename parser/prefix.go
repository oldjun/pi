package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parsePrefix() ast.Expression {
	expression := &ast.Prefix{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}
