package ast

import "github.com/oldjun/pi/token"

type String struct {
	Expression
	Token token.Token
	Value string
}
