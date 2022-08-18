package ast

import (
	"pilang/token"
)

type Assign struct {
	Expression
	Token token.Token // the = token
	Name  *Identifier
	Value Expression
}
