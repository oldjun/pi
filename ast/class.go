package ast

import (
	"pilang/token"
)

type Class struct {
	Expression
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
