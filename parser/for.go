package parser

import (
	"pilang/ast"
	"pilang/token"
)

// We first try parsing the code as a regular for loop.
// If we realize this is a for .. in we will then switch around.
func (p *Parser) parseFor() ast.Expression {
	expression := &ast.For{Token: p.currToken}
	p.nextToken()
	if !p.curTokenIs(token.IDENTIFIER) {
		return nil
	}
	if !p.peekTokenIs(token.ASSIGN) {
		return p.parseForInExpression(expression)
	}
	expression.Identifier = p.currToken.Literal
	expression.Starter = p.parseExpression(LOWEST)
	if expression.Starter == nil {
		return nil
	}
	p.nextToken()
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	expression.Condition = p.parseExpression(LOWEST)
	if expression.Condition == nil {
		return nil
	}
	p.nextToken()
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	expression.Closer = p.parseExpression(LOWEST)
	if expression.Closer == nil {
		return nil
	}
	p.nextToken()
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	if !p.curTokenIs(token.LBRACE) {
		return nil
	}
	expression.Block = p.parseBlock()
	return expression
}

// for x in [1,2,3] {}
func (p *Parser) parseForInExpression(initialExpression *ast.For) ast.Expression {
	expression := &ast.ForIn{Token: initialExpression.Token}
	if !p.curTokenIs(token.IDENTIFIER) {
		return nil
	}
	val := p.currToken.Literal
	var key string
	p.nextToken()
	if p.curTokenIs(token.COMMA) {
		p.nextToken()
		if !p.curTokenIs(token.IDENTIFIER) {
			return nil
		}
		key = val
		val = p.currToken.Literal
		p.nextToken()
	}
	expression.Key = key
	expression.Value = val
	if !p.curTokenIs(token.IN) {
		return nil
	}
	p.nextToken()
	expression.Iterable = p.parseExpression(LOWEST)
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Block = p.parseBlock()
	return expression
}
