package parser

import (
	"pilang/ast"
	"pilang/token"
	"strings"
)

func (p *Parser) parseImport() ast.Expression {
	expr := &ast.Import{Token: p.currToken}
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
	file := strings.Join(arr, ".")
	expr.File = file

	p.nextToken()
	if p.currTokenIs(token.ASTERISK) {
		expr.Everything = true
		return expr
	}

	return expr
}
