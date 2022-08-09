package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseWhile() ast.Expression {
	expression := &ast.While{Token: p.currToken}
	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Consequence = p.parseBlock()
	return expression
}
