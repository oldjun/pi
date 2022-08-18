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
