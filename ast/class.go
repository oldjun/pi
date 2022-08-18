package ast

import (
	"fmt"
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
	return fmt.Sprintf("<class:%s>", c.Name.Value)
}
