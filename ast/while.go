package ast

import (
	"bytes"
	"pilang/token"
)

type While struct {
	Token       token.Token // The 'while' token
	Condition   Expression
	Consequence *Block
}

func (w *While) expressionNode()      {}
func (w *While) TokenLiteral() string { return w.Token.Literal }
func (w *While) String() string {
	var out bytes.Buffer
	out.WriteString("while")
	out.WriteString(w.Condition.String())
	out.WriteString(" {")
	out.WriteString(w.Consequence.String())
	out.WriteString(" }")
	return out.String()
}
