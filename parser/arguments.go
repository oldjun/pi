package parser

import "github.com/oldjun/pi/ast"

func (p *Parser) parseArguments() ast.Expression {
	args := &ast.Arguments{Token: p.currToken}
	return args
}
