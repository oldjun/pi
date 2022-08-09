package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Array struct {
	token.Token // the '[' token
	Elements    []Expression
}

func (a *Array) expressionNode()      {}
func (a *Array) TokenLiteral() string { return a.Token.Literal }
func (a *Array) String() string {
	var out bytes.Buffer
	var elements []string
	for _, el := range a.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
