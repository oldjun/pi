package ast

import (
	"pilang/token"
)

type Postfix struct {
	Expression
	Token    token.Token // The operator token, e.g. ++,--
	Left     Expression
	Operator string
}
