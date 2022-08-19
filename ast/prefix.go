package ast

import (
	"github.com/oldjun/pi/token"
)

type Prefix struct {
	Expression
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}
