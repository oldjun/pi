package ast

import (
	"bytes"
	"pilang/token"
)

type IndexAssignment struct {
	Token token.Token // the '=' token
	Name  *IndexExpression
	Value Expression
}

func (ia *IndexAssignment) expressionNode()      {}
func (ia *IndexAssignment) TokenLiteral() string { return ia.Token.Literal }
func (ia *IndexAssignment) String() string {
	var out bytes.Buffer
	out.WriteString(ia.Name.String())
	out.WriteString(" = ")
	out.WriteString(ia.Value.String())
	return out.String()
}
