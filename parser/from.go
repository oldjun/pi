package parser

import (
	"pilang/ast"
	"pilang/token"
	"strings"
)

func (p *Parser) parseImport() ast.Expression {
	expr := &ast.From{Token: p.currToken}
	expr.Identifiers = make(map[string]*ast.Identifier)

	p.nextToken()
	var arr []string
	for !p.currTokenIs(token.IMPORT) {
		if p.currTokenIs(token.IDENTIFIER) {
			arr = append(arr, p.currToken.Literal)
		}
		p.nextToken()
		if p.currTokenIs(token.DOT) {
			p.nextToken()
		}
	}
	file := strings.Join(arr, "/")
	expr.File = file

	for p.peekToken.Line == p.currToken.Line {
		p.nextToken()
		if p.currTokenIs(token.ASTERISK) {
			expr.Everything = true
			return expr
		}
		identifier := &ast.Identifier{Value: p.currToken.Literal}
		alias := p.currToken.Literal
		if p.peekTokenIs(token.AS) {
			p.nextToken()
			p.nextToken()
			alias = p.currToken.Literal
		}
		expr.Identifiers[alias] = identifier
		if p.peekTokenIs(token.COMMA) {
			p.nextToken()
		}
	}
	return expr
}
