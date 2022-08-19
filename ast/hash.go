package ast

import (
	"github.com/oldjun/pi/token"
)

type Hash struct {
	Expression
	token.Token // the '{' token
	Pairs       map[Expression]Expression
}
