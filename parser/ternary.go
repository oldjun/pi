package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseTernary(condition ast.Expression) ast.Expression {
	expression := &ast.Ternary{
		Token:     p.currToken,
		Condition: condition,
	}
	p.nextToken() // skip the '?'
	precedence := p.currPrecedence()
	expression.IfTrue = p.parseExpression(precedence)
	if !p.expectPeek(token.COLON) {
		return nil
	}
	p.nextToken()
	expression.IfFalse = p.parseExpression(precedence)
	return expression
}
