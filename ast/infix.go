package ast

import (
	"bytes"
	"pilang/token"
)

type Infix struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (i *Infix) expressionNode()      {}
func (i *Infix) TokenLiteral() string { return i.Token.Literal }
func (i *Infix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")
	return out.String()
}
