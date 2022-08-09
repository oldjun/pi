package ast

import (
	"bytes"
	"pilang/token"
)

type Ternary struct {
	Token     token.Token // The ? token
	Condition Expression
	IfTrue    Expression
	IfFalse   Expression
}

func (t *Ternary) expressionNode()      {}
func (t *Ternary) TokenLiteral() string { return t.Token.Literal }
func (t *Ternary) String() string {
	var out bytes.Buffer
	out.WriteString(t.Condition.String())
	out.WriteString(" ? ")
	out.WriteString(t.IfTrue.String())
	out.WriteString(" : ")
	out.WriteString(t.IfFalse.String())
	return out.String()
}
