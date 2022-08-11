package ast

import "pilang/token"

type Null struct {
	Token token.Token // the 'null' token
}

func (b *Null) expressionNode()      {}
func (b *Null) TokenLiteral() string { return b.Token.Literal }
func (b *Null) String() string       { return "null" }
