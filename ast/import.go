package ast

import (
	"pilang/token"
)

type Import struct {
	Expression
	Token       token.Token
	Identifiers map[string]*Identifier
}
