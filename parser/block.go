package parser

import (
	"github.com/oldjun/pi/ast"
	"github.com/oldjun/pi/token"
)

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Token: p.currToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.currTokenIs(token.RBRACE) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
