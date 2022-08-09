package parser

import "pilang/ast"

func (p *Parser) parseAssignment(left ast.Expression) ast.Expression {
	switch left.(type) {
	case *ast.Identifier:
		exp := &ast.Assign{
			Token: p.currToken,
			Name:  left.(*ast.Identifier),
		}
		precedence := p.currPrecedence()
		p.nextToken()
		exp.Value = p.parseExpression(precedence)
		return exp
	case *ast.IndexExpression:
		exp := &ast.IndexAssignment{
			Token: p.currToken,
			Name:  left.(*ast.IndexExpression),
		}
		precedence := p.currPrecedence()
		p.nextToken()
		exp.Value = p.parseExpression(precedence)
		return exp
	case *ast.PropertyExpression:
		exp := &ast.PropertyAssignment{
			Token: p.currToken,
			Name:  left.(*ast.PropertyExpression),
		}
		precedence := p.currPrecedence()
		p.nextToken()
		exp.Value = p.parseExpression(precedence)
		return exp
	default:
		return nil
	}
}
