package ast

import (
	"github.com/oldjun/pi/token"
)

type Postfix struct {
	Expression
	Token    token.Token // The operator token, e.g. ++,--
	Left     Expression
	Operator string
}
