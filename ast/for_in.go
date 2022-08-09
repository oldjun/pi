package ast

import (
	"bytes"
	"pilang/token"
)

type ForIn struct {
	Token    token.Token // The 'for' token
	Key      string
	Value    string
	Iterable Expression // An expression that should return an iterable ([1, 2, 3] or x in 1..10)
	Block    *Block     // The block executed inside the for loop
}

func (f *ForIn) expressionNode()      {}
func (f *ForIn) TokenLiteral() string { return f.Token.Literal }
func (f *ForIn) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	if f.Key != "" {
		out.WriteString(f.Key + ", ")
	}
	out.WriteString(f.Value)
	out.WriteString(" in ")
	out.WriteString(f.Iterable.String())
	out.WriteString(f.Block.String())
	return out.String()
}
