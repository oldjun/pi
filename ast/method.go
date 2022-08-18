package ast

import (
	"pilang/token"
)

type Method struct {
	Expression
	Token     token.Token // The operator token .
	Object    Expression
	Method    Expression
	Arguments []Expression
}
