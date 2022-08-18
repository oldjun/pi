package parser

import "pilang/ast"

func (p *Parser) parseDecorator() ast.Expression {
	dc := &ast.Decorator{Token: p.currToken}
	defer (func() {
		p.nextToken()
		exp := p.parseExpressionStatement()
		switch fn := exp.Expression.(type) {
		case *ast.Function:
			if fn.Name == "" {
				return
			}
			dc.Decorated = fn
		case *ast.Decorator:
			dc.Decorated = fn
		default:
			return
		}
	})()

	p.nextToken()
	exp := p.parseExpressionStatement()
	dc.Function = exp.Expression
	return dc
}
