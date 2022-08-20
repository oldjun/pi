package ast

import "github.com/oldjun/pi/token"

type Async struct {
	Statement
	Token token.Token // the 'async' token
	Call  Expression
}
