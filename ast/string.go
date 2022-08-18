package ast

import "pilang/token"

type String struct {
	Expression
	Token token.Token
	Value string
}
