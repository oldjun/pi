package ast

import "github.com/oldjun/pi/token"

type ExpressionStatement struct {
	Statement
	Token      token.Token // the first token of the expression
	Expression Expression
}
