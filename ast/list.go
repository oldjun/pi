package ast

import (
	"github.com/oldjun/pi/token"
)

type List struct {
	Expression
	Token    token.Token // the '[' token
	Elements []Expression
}
