package ast

import "pilang/token"

type Decorator struct {
	Token      token.Token // @
	Expression Expression
	Decorated  Expression
}

func (d *Decorator) expressionNode()      {}
func (d *Decorator) TokenLiteral() string { return d.Token.Literal }
func (d *Decorator) String() string       { return d.Expression.String() }
