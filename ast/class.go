package ast

import (
	"bytes"
	"pilang/token"
)

type Class struct {
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}

func (c *Class) expressionNode()      {}
func (c *Class) TokenLiteral() string { return c.Token.Literal }
func (c *Class) String() string {
	var out bytes.Buffer
	return out.String()
}
