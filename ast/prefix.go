package ast

import (
	"bytes"
	"pilang/token"
)

type Prefix struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (p *Prefix) expressionNode()      {}
func (p *Prefix) TokenLiteral() string { return p.Token.Literal }
func (p *Prefix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.String())
	out.WriteString(")")
	return out.String()
}
