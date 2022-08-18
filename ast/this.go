package ast

import (
	"pilang/token"
)

type This struct {
	Token token.Token
}

func (t *This) expressionNode()      {}
func (t *This) TokenLiteral() string { return t.Token.Literal }
func (t *This) String() string       { return t.Token.Literal }
