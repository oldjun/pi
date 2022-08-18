package ast

import (
	"pilang/token"
)

type PropertyExpression struct {
	Expression
	Token    token.Token // The . token
	Object   Expression
	Property Expression
}
