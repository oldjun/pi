package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Token: p.currToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.curTokenIs(token.RBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
