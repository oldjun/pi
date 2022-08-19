package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseImport() ast.Expression {
	expr := &ast.Import{Token: p.currToken}
	expr.Identifiers = make(map[string]*ast.Identifier)
	for p.currToken.Line == p.peekToken.Line {
		p.nextToken()
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
		if p.peekTokenIs(token.EOF) {
			break
		}
	}
	return expr
}
