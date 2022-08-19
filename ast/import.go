package ast

import (
	"github.com/oldjun/pi/token"
)

type Import struct {
	Expression
	Token       token.Token
	Identifiers map[string]*Identifier
}
