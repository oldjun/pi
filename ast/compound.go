package ast

import (
	"bytes"
	"pilang/token"
)

type Compound struct {
	Token    token.Token // The operator token, e.g. +=,-=,*=,/=,%=
	Left     Expression
	Operator string
	Right    Expression
}

func (c *Compound) expressionNode()      {}
func (c *Compound) TokenLiteral() string { return c.Token.Literal }
func (c *Compound) String() string {
	var out bytes.Buffer
	out.WriteString(c.Left.String())
	out.WriteString(" " + c.Operator + " ")
	out.WriteString(c.Right.String())
	return out.String()
}
