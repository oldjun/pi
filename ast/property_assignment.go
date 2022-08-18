package ast

import (
	"pilang/token"
)

type PropertyAssignment struct {
	Expression
	Token token.Token // the '=' token
	Name  *PropertyExpression
	Value Expression
}
