package ast

import (
	"pilang/token"
)

type List struct {
	Expression
	token.Token // the '[' token
	Elements    []Expression
}
