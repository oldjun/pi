package ast

import "github.com/oldjun/pi/token"

type Identifier struct {
	Expression
	Token token.Token // the token.IDENTIFIER token
	Value string
}
