package ast

import (
	"bytes"
	"pilang/token"
)

type Postfix struct {
	Token    token.Token // The operator token, e.g. ++,--
	Left     Expression
	Operator string
}

func (p *Postfix) expressionNode()      {}
func (p *Postfix) TokenLiteral() string { return p.Token.Literal }
func (p *Postfix) String() string {
	var out bytes.Buffer
	out.WriteString(p.Token.Literal)
	out.WriteString(p.Operator)
	return out.String()
}
