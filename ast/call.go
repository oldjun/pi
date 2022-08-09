package ast

import (
	"bytes"
	"pilang/token"
	"strings"
)

type Call struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or Function
	Arguments []Expression
}

func (c *Call) expressionNode()      {}
func (c *Call) TokenLiteral() string { return c.Token.Literal }
func (c *Call) String() string {
	var out bytes.Buffer
	var args []string
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
