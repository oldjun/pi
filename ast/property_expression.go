package ast

import (
	"bytes"
	"pilang/token"
)

type PropertyExpression struct {
	Token    token.Token // The . token
	Object   Expression
	Property Expression
}

func (pe *PropertyExpression) expressionNode()      {}
func (pe *PropertyExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PropertyExpression) String() string {
	var out bytes.Buffer
	out.WriteString(pe.Object.String())
	out.WriteString(".")
	out.WriteString(pe.Property.String())
	return out.String()
}
