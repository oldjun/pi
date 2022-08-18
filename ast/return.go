package ast

import (
	"pilang/token"
)

type Return struct {
	Statement
	Token token.Token // the 'return' token
	Value Expression
}
