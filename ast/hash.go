package ast

import (
	"pilang/token"
)

type Hash struct {
	Expression
	token.Token // the '{' token
	Pairs       map[Expression]Expression
}
