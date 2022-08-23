package ast

import "github.com/oldjun/pi/token"

type Defer struct {
	Statement
	Token token.Token // the 'defer' token
	Call  Expression
}
