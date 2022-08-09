package parser

import "pilang/ast"

// @decorator
// @decorator(1, 2)
func (p *Parser) parseDecorator() ast.Expression {
	dc := &ast.Decorator{Token: p.currToken}
	// A decorator always proceeds a named function or another decorator,
	// so once we're done
	// with our "own" parsing, we can defer
	// to parsing the next statement, which
	// should either return another decorator or a function.
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
	dc.Expression = exp.Expression
	return dc
}
