package ast

import (
	"pilang/token"
)

type Super struct {
	Expression
	Token token.Token
}
