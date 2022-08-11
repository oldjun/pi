package ast

import (
	"pilang/token"
)

type Break struct {
	Token token.Token // the 'break' token
}

func (b *Break) statementNode()       {}
func (b *Break) TokenLiteral() string { return b.Token.Literal }
func (b *Break) String() string       { return "break" }
