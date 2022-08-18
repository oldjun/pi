package ast

import (
	"pilang/token"
)

type IndexAssignment struct {
	Expression
	Token token.Token // the '=' token
	Name  *IndexExpression
	Value Expression
}
