package parser

import (
	"fmt"
	"github.com/oldjun/pi/ast"
	"strconv"
)

func (p *Parser) parseInteger() ast.Expression {
	lit := &ast.Integer{Token: p.currToken}
	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
