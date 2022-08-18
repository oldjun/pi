package ast

import (
	"pilang/token"
)

type Boolean struct {
	Expression
	Token token.Token
	Value bool
}
