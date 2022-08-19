package ast

import (
	"github.com/oldjun/pi/token"
)

type List struct {
	Expression
	token.Token // the '[' token
	Elements    []Expression
}
