package ast

import (
	"pilang/token"
)

type Break struct {
	Statement
	Token token.Token // the 'break' token
}
