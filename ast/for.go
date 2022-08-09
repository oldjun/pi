package ast

import (
	"bytes"
	"pilang/token"
)

type For struct {
	Token      token.Token // The 'for' token
	Identifier string      // "i"
	Starter    Expression  // i = 0
	Closer     Expression  // i++
	Condition  Expression  // i < 1
	Block      *Block      // The block executed inside the for loop
}

func (f *For) expressionNode()      {}
func (f *For) TokenLiteral() string { return f.Token.Literal }
func (f *For) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	out.WriteString(f.Starter.String())
	out.WriteString(";")
	out.WriteString(f.Condition.String())
	out.WriteString(";")
	out.WriteString(f.Closer.String())
	out.WriteString(";")
	out.WriteString(f.Block.String())
	return out.String()
}
