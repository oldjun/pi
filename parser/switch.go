package parser

import (
	"fmt"
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseSwitch() ast.Expression {
	expr := &ast.Switch{Token: p.currToken}
	p.nextToken()
	expr.Value = p.parseExpression(LOWEST)
	if expr.Value == nil {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	p.nextToken()
	for !p.currTokenIs(token.RBRACE) {
		switchCase := &ast.Case{Token: p.currToken}
		if p.currTokenIs(token.DEFAULT) {
			switchCase.Default = true
		} else if p.currTokenIs(token.CASE) {
			p.nextToken()
			switchCase.Values = append(switchCase.Values, p.parseExpression(LOWEST))
			for p.peekTokenIs(token.COMMA) {
				p.nextToken()
				p.nextToken()
				switchCase.Values = append(switchCase.Values, p.parseExpression(LOWEST))
			}
		} else {
			// error - unexpected token
			p.errors = append(p.errors, fmt.Sprintf("expected case|default, got %s", p.currToken.Type))
			return nil
		}
		if !p.expectPeek(token.COLON) {
			return nil
		}
		switchCase.Body = &ast.Block{Token: p.currToken}
		p.nextToken()
		for !p.currTokenIs(token.CASE) && !p.currTokenIs(token.DEFAULT) {
			if p.currTokenIs(token.RBRACE) {
				break
			}
			stmt := p.parseStatement()
			switchCase.Body.Statements = append(switchCase.Body.Statements, stmt)
			p.nextToken()
		}
		expr.Cases = append(expr.Cases, switchCase)
	}
	if !p.currTokenIs(token.RBRACE) {
		return nil
	}
	return expr
}
