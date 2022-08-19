package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseIndex(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.currToken, Left: left}
	if p.peekTokenIs(token.COLON) {
		exp.Index = &ast.Integer{Token: token.Token{Type: token.INT, Literal: "0"}, Value: 0}
		exp.IsRange = true
	} else {
		p.nextToken()
		exp.Index = p.parseExpression(LOWEST)
	}
	if p.peekTokenIs(token.COLON) {
		exp.IsRange = true
		p.nextToken()
		if p.peekTokenIs(token.RBRACKET) {
			exp.End = nil
		} else {
			p.nextToken()
			exp.End = p.parseExpression(LOWEST)
		}
	}
	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	return exp
}
