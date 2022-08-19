package ast

import (
	"github.com/oldjun/pi/token"
)

type PropertyExpression struct {
	Expression
	Token    token.Token // The . token
	Object   Expression
	Property Expression
}
