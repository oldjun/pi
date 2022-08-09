package ast

import (
	"bytes"
	"pilang/token"
)

type Block struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (b *Block) statementNode()       {}
func (b *Block) TokenLiteral() string { return b.Token.Literal }
func (b *Block) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
