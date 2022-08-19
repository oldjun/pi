package ast

import (
	"github.com/oldjun/pi/token"
)

type IndexAssignment struct {
	Expression
	Token token.Token // the '=' token
	Name  *IndexExpression
	Value Expression
}
