package ast

import (
	"github.com/oldjun/pi/token"
)

type While struct {
	Expression
	Token     token.Token // The 'while' token
	Condition Expression
	Block     *Block
}
