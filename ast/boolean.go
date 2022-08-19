package ast

import (
	"github.com/oldjun/pi/token"
)

type Boolean struct {
	Expression
	Token token.Token
	Value bool
}
