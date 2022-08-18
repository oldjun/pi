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
