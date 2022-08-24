package ast

import (
	"github.com/oldjun/pi/token"
)

type Hash struct {
	Expression
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}
