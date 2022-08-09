package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseGroup() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return exp
}
