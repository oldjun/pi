package ast

import (
	"github.com/oldjun/pi/token"
)

type This struct {
	Expression
	Token token.Token
}
