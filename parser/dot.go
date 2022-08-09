package parser

import (
	"pilang/ast"
	"pilang/token"
)

// some.function() or some.property
func (p *Parser) parseDot(object ast.Expression) ast.Expression {
	tok := p.currToken
	precedence := p.currPrecedence()
	p.nextToken()

	// Here we try to figure out if we're
	// in front of a method or a property accessor.
	//
	// If the token after the identifier
	// is a (, then we're expecting this
	// to me a method (x.f()), else we'll
	// assume it's a property (x.p).
	if p.peekTokenIs(token.LPAREN) {
		exp := &ast.Method{Token: tok, Object: object}
		exp.Method = p.parseExpression(precedence)
		p.nextToken()
		exp.Arguments = p.parseExpressionList(token.RPAREN)
		return exp
	} else {
		// support assignment to hash property h.a = 1
		exp := &ast.PropertyExpression{Token: tok, Object: object}
		exp.Property = p.parseIdentifier()
		return exp
	}
}
