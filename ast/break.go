package ast

import (
	"github.com/oldjun/pi/token"
)

type Break struct {
	Statement
	Token token.Token // the 'break' token
}
