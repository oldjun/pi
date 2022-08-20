package ast

import "github.com/oldjun/pi/token"

type Arguments struct {
	Expression
	Token token.Token
}
