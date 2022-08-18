package ast

import (
	"pilang/token"
)

type Infix struct {
	Expression
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}
