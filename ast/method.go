package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Method struct {
	Token     token.Token // The operator token .
	Object    Expression
	Method    Expression
	Arguments []Expression
}

func (m *Method) expressionNode()      {}
func (m *Method) TokenLiteral() string { return m.Token.Literal }
func (m *Method) String() string {
	var out bytes.Buffer
	var args []string
	for _, a := range m.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(m.Object.String())
	out.WriteString(".")
	out.WriteString(m.Method.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
