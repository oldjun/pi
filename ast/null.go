package ast

import "github.com/oldjun/pi/token"

type Null struct {
	Expression
	Token token.Token // the 'null' token
}
