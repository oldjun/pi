package ast

import (
	"github.com/oldjun/pi/token"
)

type Method struct {
	Expression
	Token     token.Token // The operator token .
	Object    Expression
	Method    Expression
	Arguments []Expression
}
