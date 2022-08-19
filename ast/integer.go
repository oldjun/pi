package ast

import (
	"github.com/oldjun/pi/token"
)

type Integer struct {
	Expression
	Token token.Token
	Value int64
}
