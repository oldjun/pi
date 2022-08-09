package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseCall(function ast.Expression) ast.Expression {
	exp := &ast.Call{Token: p.currToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

//func (p *Parser) parseCallArguments() []ast.Expression {
//	var args []ast.Expression
//	if p.peekTokenIs(token.RPAREN) {
//		p.nextToken()
//		return args
//	}
//	p.nextToken()
//	args = append(args, p.parseExpression(LOWEST))
//	for p.peekTokenIs(token.COMMA) {
//		p.nextToken()
//		p.nextToken()
//		args = append(args, p.parseExpression(LOWEST))
//	}
//	if !p.expectPeek(token.RPAREN) {
//		return nil
//	}
//	return args
//}
