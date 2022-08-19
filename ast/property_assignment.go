package ast

import (
	"github.com/oldjun/pi/token"
)

type PropertyAssignment struct {
	Expression
	Token token.Token // the '=' token
	Name  *PropertyExpression
	Value Expression
}
