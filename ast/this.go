package ast

import (
	"pilang/token"
)

type This struct {
	Expression
	Token token.Token
}
