package ast

import "github.com/oldjun/pi/token"

type Float struct {
	Expression
	Token token.Token
	Value float64
}
