package ast

import (
	"github.com/oldjun/pi/token"
)

type Return struct {
	Statement
	Token token.Token // the 'return' token
	Value Expression
}
