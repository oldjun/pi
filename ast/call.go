package ast

import (
	"github.com/oldjun/pi/token"
)

type Call struct {
	Expression
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or Function
	Arguments []Expression
}
