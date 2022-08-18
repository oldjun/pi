package ast

import "pilang/token"

type Float struct {
	Expression
	Token token.Token
	Value float64
}
