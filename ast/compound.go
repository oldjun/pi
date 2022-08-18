package ast

import (
	"pilang/token"
)

type Compound struct {
	Expression
	Token    token.Token // The operator token, e.g. +=,-=,*=,/=,%=
	Left     Expression
	Operator string
	Right    Expression
}
