package ast

import (
	"pilang/token"
)

type Integer struct {
	Token token.Token
	Value int64
}

func (i *Integer) expressionNode()      {}
func (i *Integer) TokenLiteral() string { return i.Token.Literal }
func (i *Integer) String() string       { return i.Token.Literal }
