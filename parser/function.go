package parser

import (
	"pilang/ast"
	"pilang/token"
)

func (p *Parser) parseFunction() ast.Expression {
	fn := &ast.Function{Token: p.currToken}
	if p.peekTokenIs(token.IDENTIFIER) {
		p.nextToken()
		fn.Name = p.currToken.Literal
	}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	if !p.parseFunctionParameters(fn) {
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	fn.Body = p.parseBlock()
	return fn
}

func (p *Parser) parseFunctionParameters(fn *ast.Function) bool {
	fn.Defaults = make(map[string]ast.Expression)
	for !p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		if p.currTokenIs(token.COMMA) {
			continue
		}
		if p.currTokenIs(token.ASTERISK) {
			p.nextToken()
			fn.Args = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			if !p.peekTokenIs(token.COMMA) {
				break
			}
			p.nextToken()
			if !p.peekTokenIs(token.ASTERISK_ASTERISK) {
				break
			}
			p.nextToken()
			p.nextToken()
			fn.KwArgs = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			break
		} else if p.currTokenIs(token.ASTERISK_ASTERISK) {
			p.nextToken()
			fn.KwArgs = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			break
		} else {
			ident := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			fn.Parameters = append(fn.Parameters, ident)
			if p.peekTokenIs(token.ASSIGN) {
				p.nextToken()
				p.nextToken()
				fn.Defaults[ident.String()] = p.parseExpression(LOWEST)
			} else {
				if len(fn.Defaults) > 0 {
					return false
				}
			}
		}
	}
	if !p.expectPeek(token.RPAREN) {
		return false
	}
	return true
}
