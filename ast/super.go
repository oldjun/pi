package ast

import (
	"github.com/oldjun/pi/token"
)

type Super struct {
	Expression
	Token token.Token
}
