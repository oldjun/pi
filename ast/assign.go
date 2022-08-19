package ast

import (
	"github.com/oldjun/pi/token"
)

type Assign struct {
	Expression
	Token token.Token // the = token
	Name  *Identifier
	Value Expression
}
