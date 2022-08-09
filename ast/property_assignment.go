package ast

import (
	"bytes"
	"pilang/token"
)

type PropertyAssignment struct {
	Token token.Token // the '=' token
	Name  *PropertyExpression
	Value Expression
}

func (pa *PropertyAssignment) expressionNode()      {}
func (pa *PropertyAssignment) TokenLiteral() string { return pa.Token.Literal }
func (pa *PropertyAssignment) String() string {
	var out bytes.Buffer
	out.WriteString(pa.Name.String())
	out.WriteString(" = ")
	out.WriteString(pa.Value.String())
	return out.String()
}
