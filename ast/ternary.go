package ast

import (
	"github.com/oldjun/pi/token"
)

type Ternary struct {
	Expression
	Token     token.Token // The ? token
	Condition Expression
	IfTrue    Expression
	IfFalse   Expression
}
