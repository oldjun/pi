package parser

import (
	"fmt"
	"github.com/oldjun/pi/ast"
	"strconv"
)

func (p *Parser) parseFloat() ast.Expression {
	lit := &ast.Float{Token: p.currToken}
	value, err := strconv.ParseFloat(p.currToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
