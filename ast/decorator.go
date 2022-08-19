package ast

import "github.com/oldjun/pi/token"

type Decorator struct {
	Expression
	Token     token.Token // @
	Function  Expression
	Decorated Expression
}
