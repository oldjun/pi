package ast

import (
	"bytes"
	"pilang/token"
)

type Scenario struct {
	Condition   Expression
	Consequence *Block
}

type If struct {
	Token     token.Token // The 'if' token
	Scenarios []*Scenario
}

func (i *If) expressionNode()      {}
func (i *If) TokenLiteral() string { return i.Token.Literal }
func (i *If) String() string {
	var out bytes.Buffer
	for i, s := range i.Scenarios {
		if i != 0 {
			out.WriteString("else")
			out.WriteString(" ")
		}
		out.WriteString("if")
		out.WriteString(s.Condition.String())
		out.WriteString(" ")
		out.WriteString(s.Consequence.String())
	}
	return out.String()
}
