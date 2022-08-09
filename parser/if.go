package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseIf() ast.Expression {
	expression := &ast.If{Token: p.currToken}
	var scenarios []*ast.Scenario

	p.nextToken()
	scenario := &ast.Scenario{}
	scenario.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	scenario.Consequence = p.parseBlock()
	scenarios = append(scenarios, scenario)
	// If we encounter ELSEs then let's add more
	// scenarios to our expression.
	for p.peekTokenIs(token.ELSE) {
		p.nextToken()
		p.nextToken()
		scenario := &ast.Scenario{}
		// ELSE IF
		if p.curTokenIs(token.IF) {
			p.nextToken()
			scenario.Condition = p.parseExpression(LOWEST)
			if !p.expectPeek(token.LBRACE) {
				return nil
			}
		} else {
			tok := &token.Token{Type: token.LookupIdent("true"), Literal: "true"}
			scenario.Condition = &ast.Boolean{Token: *tok, Value: true}
		}
		scenario.Consequence = p.parseBlock()
		scenarios = append(scenarios, scenario)
	}
	expression.Scenarios = scenarios
	return expression
}
