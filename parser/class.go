package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseClass() ast.Expression {
	expression := &ast.Class{Token: p.currToken}
	p.nextToken()
	expression.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	if p.peekTokenIs(token.COLON) {
		p.nextToken()
		if !p.expectPeek(token.IDENTIFIER) {
			return nil
		}
		expression.Super = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Body = p.parseBlock()
	return expression
}
