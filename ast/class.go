package ast

import (
	"github.com/oldjun/pi/token"
)

type Class struct {
	Expression
	Token token.Token
	Name  *Identifier
	Super *Identifier
	Body  *Block
}
