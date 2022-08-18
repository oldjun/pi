package ast

import (
	"pilang/token"
)

type Integer struct {
	Expression
	Token token.Token
	Value int64
}
