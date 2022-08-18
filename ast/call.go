package ast

import (
	"pilang/token"
)

type Call struct {
	Expression
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or Function
	Arguments []Expression
}
