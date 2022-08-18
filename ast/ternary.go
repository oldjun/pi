package ast

import (
	"pilang/token"
)

type Ternary struct {
	Expression
	Token     token.Token // The ? token
	Condition Expression
	IfTrue    Expression
	IfFalse   Expression
}
