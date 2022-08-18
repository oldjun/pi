package ast

import (
	"pilang/token"
)

type IndexExpression struct {
	Expression
	Token   token.Token // The [ token
	Left    Expression
	Index   Expression
	IsRange bool       // whether the expression is a range [1:10]
	End     Expression // the end of the range, if the expression is a range
}
