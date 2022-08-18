package ast

import (
	"pilang/token"
)

type Prefix struct {
	Expression
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}
