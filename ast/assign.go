package ast

import (
	"bytes"
	"pilang/token"
)

type Assign struct {
	Token token.Token // the = token
	Name  *Identifier
	Value Expression
}

func (a *Assign) expressionNode()      {}
func (a *Assign) TokenLiteral() string { return a.Token.Literal }
func (a *Assign) String() string {
	var out bytes.Buffer
	out.WriteString(a.Name.String())
	out.WriteString(" = ")
	out.WriteString(a.Value.String())
	return out.String()
}
