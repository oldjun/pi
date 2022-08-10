package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type List struct {
	token.Token // the '[' token
	Elements    []Expression
}

func (l *List) expressionNode()      {}
func (l *List) TokenLiteral() string { return l.Token.Literal }
func (l *List) String() string {
	var out bytes.Buffer
	var elements []string
	for _, el := range l.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
