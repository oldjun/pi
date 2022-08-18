package ast

import "pilang/token"

type Decorator struct {
	Expression
	Token     token.Token // @
	Function  Expression
	Decorated Expression
}
